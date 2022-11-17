package http

import (
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/command"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/ddd"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/discovery"
	"github.com/itsfunny/go-cell/di"
	"go.uber.org/fx"
)

var (
	HttpDiscoveryModule di.OptionBuilder = func() fx.Option {
		return fx.Options(
			ddd.DIHttpEnvelopeHandler,
			command.Commands,
			di.RegisterExtension(discovery.NewHttpDiscoveryComponent),
		)
	}
)
