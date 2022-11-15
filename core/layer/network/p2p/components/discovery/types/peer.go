package types

type PeerId string

type IPeerManager interface {
	GetMembership() map[PeerId]IPeerNode
	GetSelfNode() IPeerNode
	Register(wrapper PeerWrapper)
	Have(node PeerId) bool
	GetByPeerId(id PeerId) IPeerNode
}

type IPeerNode interface {
	PeerId() PeerId
	MetaData() PeerMetaData
}

type PeerWrapper struct {
	PeerId   PeerId
	MetaData PeerMetaData
}

func (p PeerId) ToString() string {
	return string(p)
}
