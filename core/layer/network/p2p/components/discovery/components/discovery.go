package components

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/enums"
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/config"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/codec"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
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

	pingPongEnvelopeCreateF func() *types2.Envelope
}

func NewBaseDiscoveryComponent(
	ddd *component.DDDComponent, cdc *codec.CodecComponent,
	peerManager types.IPeerManager,
	PingPongEnvelopeCreateF func() *types2.Envelope,
	internal types.DiscoveryComponent,
) *BaseDiscoveryComponent {
	ret := &BaseDiscoveryComponent{PeerManager: peerManager, internal: internal}
	ret.BaseComponent = component.NewBaseComponent(enums.DiscoveryModule, internal, ddd, cdc)
	ret.pingPongEnvelopeCreateF = PingPongEnvelopeCreateF
	return ret
}

func (b BaseDiscoveryComponent) OnStart(ctx *services.StartCTX) error {
	go b.pingPong()
	return nil
}

func (b BaseDiscoveryComponent) pingPong() {
	timer := time.NewTimer(time.Second * time.Duration(b.Config.PingPongPeriod))
	for {
		select {
		case <-b.Quit():
			return
		case <-timer.C:
			ctx := b.GetContext()
			cellCtx := sdk.EmptyCellContext(ctx)
			b.BroadCast(cellCtx, types.BroadCastRequest{
				Envelop: b.pingPongEnvelopeCreateF(),
			})
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
			To:      mem,
			Envelop: req.Envelop,
		})
		if nil != err {
			b.Logger.Error("send to peer failed", "err", err.Error(), "id", id)
		}
		promises = append(promises, p)
	}

	ret := types.BroadCastResponse{}

	return ret
}
