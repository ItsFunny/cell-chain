package ddd

import "github.com/itsfunny/cell-chain/common/component"

var DIHttpEnvelopeHandler = component.RegisterEnvelopHandler(NewHttpSendPeerDDDHandler)
