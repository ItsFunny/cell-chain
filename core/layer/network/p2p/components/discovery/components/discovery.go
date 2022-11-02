package components

import (
	"github.com/itsfunny/cell-chain/common/component"
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/promise"
)

var (
	_ types.DiscoveryComponent = (*BaseDiscoveryComponent)(nil)
)

type BaseDiscoveryComponent struct {
	*component.BaseComponent
	peerManager types.IPeerManager
	internal    types.DiscoveryComponent
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
	mems := b.peerManager.GetMembership()
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
