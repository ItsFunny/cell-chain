package discovery

import (
	"context"
	"github.com/itsfunny/cell-chain/common/component"
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components"
	types2 "github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/component/codec"
)

var (
	_ types.DiscoveryComponent = (*HttpDiscoveryComponent)(nil)
)

type HttpDiscoveryComponent struct {
	*components.BaseDiscoveryComponent
}

// TODO, wrapped with fx
//func (b *HttpDiscoveryComponent) Handlers() []component.EnvelopeHandler {
//	return
//}

func NewHttpDiscoveryComponent(ctx context.Context,
	ddd *component.DDDComponent, cdc *codec.CodecComponent,
	peerManager types.IPeerManager,
	bus eventbus.ICommonEventBus,
) types.DiscoveryComponent {
	ret := &HttpDiscoveryComponent{}
	ret.BaseDiscoveryComponent = components.NewBaseDiscoveryComponent(ctx, ddd, cdc, peerManager, bus, ret)
	return ret
}

func (b *HttpDiscoveryComponent) SendToPeerAsync(ctx sdk.CellContext, req types.SendToPeerRequest) (*promise.Promise, error) {
	return b.SendAsync(ctx, func() sdk.CellRequest {
		ret := types2.HttpSendToPeerRequest{}
		return ret.From(req)
	})
}
