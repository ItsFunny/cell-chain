package types

type PeerId interface {
}
type IPeerManager interface {
	GetMembership() map[PeerId]IPeerNode
}

type IPeerNode interface {
	PeerId() PeerId
	MetaData() PeerMetaData
}
