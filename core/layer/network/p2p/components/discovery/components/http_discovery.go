package components

import (
	"github.com/itsfunny/cell-chain/common/component"
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	types2 "github.com/itsfunny/cell-chain/core/layer/network/types"
)

var (
	_ types.DiscoveryComponent = (*HttpDiscoveryComponent)(nil)
)

type HttpDiscoveryComponent struct {
	*component.BaseComponent
}

func (h *HttpDiscoveryComponent) BroadCast(ctx sdk.CellContext, req types2.BroadCastRequest) (types2.BroadCastResponse, error) {

	return types2.BroadCastResponse{}, nil
}
