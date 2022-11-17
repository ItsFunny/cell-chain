package peermanager

import (
	"github.com/itsfunny/go-cell/di"
	"go.uber.org/fx"
)

var DefaultPeerManagerModule di.OptionBuilder = func() fx.Option {
	return fx.Options(
		di.RegisterExtension(NewDefaultPeerManager),
	)
}
