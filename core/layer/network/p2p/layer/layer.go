package layer

import (
	"github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network"
	types2 "github.com/itsfunny/cell-chain/core/layer/network/types"
)

type P2PLayer interface {
	network.NetworkLayer

	BroadCast(ctx types.CellContext, req types2.BroadCastRequest) (types2.BroadCastResponse, error)
}
