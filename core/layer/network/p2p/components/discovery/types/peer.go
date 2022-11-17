package types

type PeerId string

type IPeerManager interface {
	GetMembership() map[PeerId]IPeerNode
	GetSelfNode() IPeerNode
	Register(wrapper *PeerWrapper)
	Have(node PeerId) bool
	GetByPeerId(id PeerId) IPeerNode
	ForwardMessage() chan<- interface{}
}

type IPeerNode interface {
	PeerId() PeerId
	MetaData() PeerMetaData
}

type PeerWrapper struct {
	PeerId   PeerId
	MetaData PeerMetaData
}

func NewPeerWrapper(peerId PeerId, metaData PeerMetaData) *PeerWrapper {
	return &PeerWrapper{PeerId: peerId, MetaData: metaData}
}

func (p PeerId) ToString() string {
	return string(p)
}
