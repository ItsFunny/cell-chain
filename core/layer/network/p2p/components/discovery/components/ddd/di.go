package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/go-cell/di"
	"go.uber.org/fx"
)

var (
	DiscoveryEnvelopeHandlerModule di.OptionBuilder = func() fx.Option {
		return fx.Options(
			component.RegisterEnvelopHandler(NewNewMemberHandler),
			component.RegisterEnvelopHandler(NewPingHandler),
			component.RegisterEnvelopHandler(NewPongHandler),
			component.RegisterEnvelopHandler(NewProbeEnvelopeHandler),
			component.RegisterEnvelopHandler(NewProbeSyncHandler),
			component.RegisterEnvelopHandler(NewMemberShareHandler),
		)
	}
)
