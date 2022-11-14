package types

import (
	"github.com/itsfunny/go-cell/base/core/eventbus"
)

var (
	DiscoveryEventTypeKey = "discovery.event"
	DiscoveryEvent        = "discovery"
)

func PublishSendMessageEvent(bus eventbus.ICommonEventBus, protocol string, detailMsg interface{}) {

}
