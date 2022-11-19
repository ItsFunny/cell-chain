package peermanager

import "github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"

var (
	_ types.IPeerNode = (*DefaultPeerNode)(nil)
)

type DefaultPeerNode struct {
	peerId   types.PeerId
	metaData types.PeerMetaData
}

func NewDefaultPeerNode(peerId types.PeerId, metaData types.PeerMetaData) *DefaultPeerNode {
	return &DefaultPeerNode{peerId: peerId, metaData: metaData}
}

func (d *DefaultPeerNode) PeerId() types.PeerId {
	return d.peerId
}

func (d *DefaultPeerNode) MetaData() types.PeerMetaData {
	return d.metaData
}

func (d *DefaultPeerNode) From(wp *types.PeerWrapper) {
	d.peerId = wp.PeerId
	d.metaData = wp.MetaData
}
