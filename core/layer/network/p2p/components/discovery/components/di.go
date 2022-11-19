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
		)
	}
)

type BaseDiscoveryExtension struct {
	*extension.BaseExtension
	D        types.DiscoveryComponent
	concrete extension.INodeExtension
}

func NewBaseDiscoveryExtension(d types.DiscoveryComponent, ext extension.INodeExtension) *BaseDiscoveryExtension {
	ret := &BaseDiscoveryExtension{}
	ret.BaseExtension = extension.NewBaseExtension(ret)
	ret.D = d
	ret.concrete = ext
	return ret
}

func (d *BaseDiscoveryExtension) OnExtensionInit(ctx extension.INodeContext) error {
	return d.concrete.OnExtensionInit(ctx)
}

func (d *BaseDiscoveryExtension) OnExtensionStart(ctx extension.INodeContext) error {
	return d.D.BStart(services.AsyncStartOpt)
}
