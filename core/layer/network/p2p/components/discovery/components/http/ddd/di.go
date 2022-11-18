package ddd

import "github.com/itsfunny/cell-chain/common/component"

var DIHttpDDDHandler = component.RegisterDDDHandler(NewHttpSendPeerDDDHandler)
