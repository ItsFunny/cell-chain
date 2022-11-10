package types

type PeerId string

type IPeerManager interface {
	GetMembership() map[PeerId]IPeerNode
	GetSelfNode() IPeerNode
	Register(wrapper PeerWrapper)
}

type IPeerNode interface {
	PeerId() PeerId
	MetaData() PeerMetaData
}

type PeerWrapper struct {
	PeerId   PeerId
	MetaData PeerMetaData
}
