package config

type DiscoveryConfiguration struct {
	PingPeriod   uint
	MemberPeriod uint
}

func DefaultDiscoveryConfiguration() *DiscoveryConfiguration {
	return &DiscoveryConfiguration{
		PingPeriod:   10,
		MemberPeriod: 10,
	}
}
