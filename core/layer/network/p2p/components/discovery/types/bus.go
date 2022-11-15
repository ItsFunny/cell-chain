package types

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/eventbus"
)

var (
	DiscoveryEventTypeKey = "discovery.event"
	DiscoveryEvent        = "discovery"
)

func PublishDiscoverySendMessageEvent(bus eventbus.ICommonEventBus, detailMsg interface{}) {
	bus.PublishWithEvents(context.Background(), detailMsg, map[string][]string{
		DiscoveryEventTypeKey: {DiscoveryEvent},
	})
}
