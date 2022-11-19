package http

import (
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/command"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/ddd"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/discovery"
	"github.com/itsfunny/go-cell/di"
	"github.com/itsfunny/go-cell/extension/http"
	"go.uber.org/fx"
)

var (
	HttpDiscoveryModule di.OptionBuilder = func() fx.Option {
		return fx.Options(
			http.HttpModule(),
			components.DIDiscoveryModule(),
			ddd.DIHttpDDDHandler,
			command.Commands,
			fx.Provide(discovery.NewHttpDiscoveryComponent),
		)
	}
)
