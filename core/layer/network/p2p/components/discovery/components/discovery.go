package components

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/enums"
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/config"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/codec"
	"time"
)

var (
	_ types.DiscoveryComponent = (*BaseDiscoveryComponent)(nil)
)

type BaseDiscoveryComponent struct {
	*component.BaseComponent
	PeerManager types.IPeerManager
	internal    types.DiscoveryComponent

	Config *config.DiscoveryConfiguration
	Bus    eventbus.ICommonEventBus
	sub    eventbus.Subscription
}

func NewBaseDiscoveryComponent(
	ddd *component.DDDComponent, cdc *codec.CodecComponent,
	peerManager types.IPeerManager,
	internal types.DiscoveryComponent,
) *BaseDiscoveryComponent {
	ret := &BaseDiscoveryComponent{PeerManager: peerManager, internal: internal}
	ret.BaseComponent = component.NewBaseComponent(enums.DiscoveryModule, internal, ddd, cdc)
	subscribe, err := ret.Bus.Subscribe(ret.GetContext(), "discovery", eventbus.QueryForEvent(types.DiscoveryEventTypeKey, types.DiscoveryEvent), 10)
	if nil != err {
		panic(err)
	}
	ret.sub = subscribe
	// TODO, configurable
	ret.Config = config.DefaultDiscoveryConfiguration()
	return ret
}

func (b BaseDiscoveryComponent) OnStart(ctx *services.StartCTX) error {
	go b.periodPing()
	go b.periodBroadCastMembers()
	go b.onRecv()
	return nil
}

func (b BaseDiscoveryComponent) onRecv() {
	ch := b.sub.Out()
	for {
		select {
		case msg := <-ch:
			if err := b.handleMsg(msg); nil != err {
				b.Logger.Error("handle failed", "err", err)
			}
		case <-b.sub.Canceled():
			return
		}
	}
}
func (b BaseDiscoveryComponent) handleMsg(msg interface{}) error {
	switch v := msg.(type) {
	case *types.SendToPeerRequest:
		b.SendToPeerAsync(sdk.EmptyCellContext(b.GetContext()), *v)
	default:
		b.Logger.Warn("未知的msg", "msg", msg)
	}
	return nil
}
func (b BaseDiscoveryComponent) periodPing() {
	timer := time.NewTimer(time.Second * time.Duration(b.Config.PingPeriod))
	for {
		select {
		case <-b.Quit():
			return
		case <-timer.C:
			ctx := b.GetContext()
			cellCtx := sdk.EmptyCellContext(ctx)
			selfNode := b.PeerManager.GetSelfNode()
			b.BroadCast(cellCtx, types.BroadCastRequest{
				Envelop: types.CreatePingEnvelopeRequest(b.GetCodec(),
					selfNode.PeerId(),
					selfNode.MetaData().GetOutPutAddress()),
			})
		}
	}
}

func (b BaseDiscoveryComponent) periodBroadCastMembers() {
	timer := time.NewTimer(time.Second * time.Duration(b.Config.MemberPeriod))
	for {
		select {
		case <-b.Quit():
			return
		case <-timer.C:
			selfNode := b.PeerManager.GetSelfNode()
			members := b.PeerManager.GetMembership()
			ctx := b.GetContext()
			cellCtx := sdk.EmptyCellContext(ctx)
			b.BroadCast(cellCtx, types.BroadCastRequest{
				Envelop: types.CreateMemberShareEnvelopeRequest(b.GetCodec(),
					selfNode.PeerId(), members)},
			)
		}
	}
}

func (b BaseDiscoveryComponent) SendToPeerAsync(ctx sdk.CellContext, req types.SendToPeerRequest) (*promise.Promise, error) {
	return b.internal.SendToPeerAsync(ctx, req)
}

func (b BaseDiscoveryComponent) SendToPeer(ctx sdk.CellContext, req types.SendToPeerRequest) (types.SendToPeerResponse, error) {
	async, err := b.internal.SendToPeerAsync(ctx, req)
	if nil != err {
		return types.SendToPeerResponse{}, err
	}
	ret, err := async.GetForever()
	return ret.(types.SendToPeerResponse), err
}

func (b BaseDiscoveryComponent) BroadCast(ctx sdk.CellContext, req types.BroadCastRequest) types.BroadCastResponse {
	mems := b.PeerManager.GetMembership()
	promises := make([]*promise.Promise, 0)
	for id, mem := range mems {
		p, err := b.SendToPeerAsync(ctx, types.SendToPeerRequest{
			To:      mem.MetaData().GetOutPutAddress(),
			Envelop: req.Envelop,
		})
		if nil != err {
			b.Logger.Error("send to peer failed", "err", err.Error(), "id", id)
			continue
		}
		promises = append(promises, p)
	}

	ret := types.BroadCastResponse{}

	return ret
}
