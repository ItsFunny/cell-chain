package types

type PingRequest struct {
	FromPeerId     PeerId
	FromOutPutAddr string
}

func NewPingRequest(fromPeerId PeerId, fromOutPutAddr string) *PingRequest {
	return &PingRequest{FromPeerId: fromPeerId, FromOutPutAddr: fromOutPutAddr}
}

type MembersShareRequest struct {
	FromPeerId PeerId
	KnownPeers map[PeerId]PeerMetaData
}

func NewMembersShareRequest(fromPeerId PeerId, knownPeers map[PeerId]PeerMetaData) *MembersShareRequest {
	return &MembersShareRequest{FromPeerId: fromPeerId, KnownPeers: knownPeers}
}
