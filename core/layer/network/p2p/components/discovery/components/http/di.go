package http

import (
	"errors"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/command"
	config2 "github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/config"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/ddd"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/discovery"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/peermanager"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/node/core/extension"
	"github.com/itsfunny/go-cell/component/codec"
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/extension/http"
	"github.com/itsfunny/go-cell/framework/http/config"
	"go.uber.org/fx"
	"reflect"
)

var (
	HttpDiscoveryModule di.OptionBuilder = func() fx.Option {
		return fx.Options(
			http.HttpModule(),
			components.DIDiscoveryModule(),
			ddd.DIHttpDDDHandler,
			command.Commands,
			fx.Provide(discovery.NewHttpDiscoveryComponent),
			di.RegisterExtension(NewHttpDiscoveryExtension),
		)
	}

	httpDiscoveryConfigModule = "httpDiscovery"
)

type HttpDiscoveryExtension struct {
	*components.BaseDiscoveryExtension
	cfg *config2.HttpDiscoveryConfiguration
}

func NewHttpDiscoveryExtension(d types.DiscoveryComponent) extension.INodeExtension {
	ret := &HttpDiscoveryExtension{}
	ret.BaseDiscoveryExtension = components.NewBaseDiscoveryExtension(d, ret)
	return ret
}

func (d *HttpDiscoveryExtension) OnExtensionInit(ctx extension.INodeContext) error {
	v := ctx.SwitchTo(reflect.TypeOf(&http.HttpFarmeWorkExtension{}))
	cdc := ctx.GetCodec()
	data := v.CurrentGenesis(cdc)
	ra := config.HttpConfiguration{}
	cdc.MustUnMarshal(data, &ra)
	ip := ctx.GetIp()
	node := peermanager.NewDefaultPeerNode(types.PeerId(d.cfg.PeerId), *types.NewPeerMetaData(d.cfg.OutputAddress, ip, ra.Protocol, ra.Port))
	mana := d.BaseDiscoveryExtension.D.GetPeerManager()
	mana.SetupSelfNode(node)
	return nil
}

func (b *HttpDiscoveryExtension) ConfigModuleName() string {
	return httpDiscoveryConfigModule
}

func (b *HttpDiscoveryExtension) LoadGenesis(cdc *codec.CodecComponent, data []byte) error {
	if len(data) == 0 {
		return errors.New("asd")
	}
	var cfg config2.HttpDiscoveryConfiguration
	if err := cdc.UnMarshal(data, &cfg); nil != err {
		return err
	}
	b.cfg = &cfg
	return nil
}

func (b *HttpDiscoveryExtension) CurrentGenesis(cdc *codec.CodecComponent) []byte {
	return cdc.MustMarshal(b.cfg)
}
