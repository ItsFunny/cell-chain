package enums

import logsdk "github.com/itsfunny/go-cell/sdk/log"

var (
	DDDModule          = logsdk.NewModule("DDD", 1)
	PeerManagerModule  = logsdk.NewModule("PeerManager", 2)
	DiscoveryModule    = logsdk.NewModule("Discovery", 3)
	EnvelopeDispatcher = logsdk.NewModule("envelope_dispatcher", 4)
)
