package components

import (
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/ddd"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/peermanager"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/itsfunny/go-cell/di"
	"go.uber.org/fx"
)

var (
	DIDiscoveryModule di.OptionBuilder = func() fx.Option {
		return fx.Options(
			ddd.DiscoveryEnvelopeHandlerModule,
			peermanager.DefaultPeerManagerModule,
			di.RegisterExtension(NewDiscoveryExtension),
		)
	}
)

type DiscoveryExtension struct {
	*extension.BaseExtension
	d types.DiscoveryComponent
}

func NewDiscoveryExtension(d types.DiscoveryComponent) extension.INodeExtension {
	ret := &DiscoveryExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	ret.d = d
	return ret
}

func (d *DiscoveryExtension) OnExtensionStart(ctx extension.INodeContext) error {
	return d.d.BStart(services.AsyncStartOpt)
}
