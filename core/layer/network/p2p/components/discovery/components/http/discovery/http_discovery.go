package discovery

import (
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components"
	types2 "github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/promise"
)

var (
	_ types.DiscoveryComponent = (*HttpDiscoveryComponent)(nil)
)

type HttpDiscoveryComponent struct {
	*components.BaseDiscoveryComponent
}

func (b *HttpDiscoveryComponent) SendToPeerAsync(ctx sdk.CellContext, req types.SendToPeerRequest) (*promise.Promise, error) {
	return b.SendAsync(ctx, func() sdk.CellRequest {
		ret := types2.HttpSendToPeerRequest{}
		return ret.From(req)
	})
}
