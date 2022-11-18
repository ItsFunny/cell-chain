package component

import (
	"go.uber.org/fx"
)

type DDDHandlerHolder struct {
	fx.In
	Handlers []DDDHandler `group:"ddd"`
}

type EnvelopeHandlerHolder struct {
	fx.In
	Envelopes []EnvelopeHandler `group:"envelopeHandler"`
}
