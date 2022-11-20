package types

import (
	"context"
	"github.com/itsfunny/go-cell/base/core/eventbus"
)

var (
	DiscoveryEventTypeKey = "discovery.event"
	DiscoveryEvent        = "discovery"

	PeerManagerEventTypeKey = "peerManager.event"
	peerManagerCommonEvent  = "common"
)

func PublishDiscoverySendMessageEvent(bus eventbus.ICommonEventBus, detailMsg interface{}) {
	bus.PublishWithEvents(context.Background(), detailMsg, map[string][]string{
		DiscoveryEventTypeKey: {DiscoveryEvent},
	})
}

func PublishPeerManagerEvent(bus eventbus.ICommonEventBus, msg interface{}) {
	bus.PublishWithEvents(context.Background(), msg, map[string][]string{
		PeerManagerEventTypeKey: {peerManagerCommonEvent},
	})
}

func SubscribePeerManagerCommonEvent(ctx context.Context, clientId string, bus eventbus.ICommonEventBus, cap int) eventbus.Subscription {
	sub, err := bus.Subscribe(ctx, clientId, eventbus.QueryForEvent(PeerManagerEventTypeKey, peerManagerCommonEvent), cap)
	if nil != err {
		panic(err)
	}
	return sub
}
