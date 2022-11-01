package types

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/types"
	types2 "github.com/itsfunny/cell-chain/core/layer/network/types"
)

type DiscoveryComponent interface {
	component.CellComponent
	BroadCast(ctx types.CellContext, req types2.BroadCastRequest) (types2.BroadCastResponse, error)
}
