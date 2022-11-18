package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"go.uber.org/fx"
)

var (
	DiscoveryEnvelopeHandlerModule = fx.Options(
		component.RegisterEnvelopHandler(NewNewMemberHandler),
		component.RegisterEnvelopHandler(NewPingHandler),
		component.RegisterEnvelopHandler(NewPongHandler),
		component.RegisterEnvelopHandler(NewProbeEnvelopeHandler),
		component.RegisterEnvelopHandler(NewProbeSyncHandler),
		component.RegisterEnvelopHandler(NewMemberShareHandler),
	)
)
