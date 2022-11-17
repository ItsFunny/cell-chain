package command

import (
	"github.com/itsfunny/go-cell/di"
)

var Commands = di.RegisterCommandConstructor(
	newPeer2PeerCommand,
)
