package network

import (
	"github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/common"
	nettypes "github.com/itsfunny/cell-chain/core/layer/network/types"
)

type NetworkLayer interface {
	common.Layer
	SendRequest(ctx types.CellContext, req nettypes.P2PRequest) (nettypes.P2PResponse, error)
}
