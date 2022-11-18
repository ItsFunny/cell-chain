package component

import (
	"context"
	"github.com/itsfunny/go-cell/component/routine"
	v2 "github.com/itsfunny/go-cell/component/routine/v2"
	"github.com/itsfunny/go-cell/di"
	"go.uber.org/fx"
)

const (
	envelopeHandler = "envelopeHandler"
	ddd             = "ddd"
)

var (
	DIDDDModule di.OptionBuilder = func() fx.Option {
		return fx.Options(
			di.RegisterExtension(NewDDDExtension),
			fx.Provide(NewDDDComponent),
		)
	}
	//DIEnvelopeModule di.OptionBuilder = func() fx.Option {
	//	return fx.Options()
	//}

	DIDefaultRoutineModule di.OptionBuilder = func() fx.Option {
		return fx.Provide(newV2RoutinePool)
	}
)

func newV2RoutinePool(ctx context.Context) routine.IRoutineComponent {
	return v2.NewV2RoutinePoolExecutorComponent(ctx)
}
func RegisterEnvelopHandler(constructor interface{}) fx.Option {
	return fx.Provide(fx.Annotated{
		Group:  envelopeHandler,
		Target: constructor,
	})
}

func RegisterDDDHandler(constructor interface{}) fx.Option {
	return fx.Options(
		fx.Provide(fx.Annotated{
			Group:  ddd,
			Target: constructor,
		}),
	)
}
