package dispatcher

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/go-cell/di"
	"go.uber.org/fx"
)

var (
	DIMsgDispatcherModule di.OptionBuilder = func() fx.Option {
		return fx.Options(
			component.RegisterDDDHandler(NewMsgDispatcher),
		)
	}
)
