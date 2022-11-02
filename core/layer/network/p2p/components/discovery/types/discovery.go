package types

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/go-cell/base/core/promise"
)

type DiscoveryComponent interface {
	component.CellComponent
	BroadCast(ctx types.CellContext, req BroadCastRequest) BroadCastResponse
	SendToPeer(ctx types.CellContext, req SendToPeerRequest) (SendToPeerResponse, error)
	SendToPeerAsync(ctx types.CellContext, req SendToPeerRequest) (*promise.Promise, error)
}
