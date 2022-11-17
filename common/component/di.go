package component

import "go.uber.org/fx"

const (
	envelopeHandler = "envelopeHandler"
)

func RegisterEnvelopHandler(constructor interface{}) fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  envelopeHandler,
		Target: constructor,
	})
}
