package enums

import logsdk "github.com/itsfunny/go-cell/sdk/log"

var (
	DDDModule          = logsdk.NewModule("DDD", 1)
	PeerManagerModule  = logsdk.NewModule("PeerManager", 2)
	DiscoveryModule    = logsdk.NewModule("Discovery", 3)
	EnvelopeDispatcher = logsdk.NewModule("envelope_dispatcher", 4)

	PingPongHandler    = logsdk.NewModule("ping_pong", 5)
	MemberShareHandler = logsdk.NewModule("member_share", 6)
	ProbeHandler       = logsdk.NewModule("probe", 7)
)
