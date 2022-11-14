package types

import "fmt"

type PingRequest struct {
	FromPeerId     PeerId
	FromOutPutAddr string
}

func NewPingRequest(fromPeerId PeerId, fromOutPutAddr string) *PingRequest {
	return &PingRequest{FromPeerId: fromPeerId, FromOutPutAddr: fromOutPutAddr}
}

type MembersShareRequest struct {
	FromPeerId PeerId
	KnownPeers map[PeerId]string
}

func (m MembersShareRequest) String() string {
	return fmt.Sprintf("from:%s,knownPeers:%v", m.FromPeerId, m.KnownPeers)
}

func NewMembersShareRequest(fromPeerId PeerId, knownPeers map[PeerId]string) *MembersShareRequest {
	return &MembersShareRequest{FromPeerId: fromPeerId, KnownPeers: knownPeers}
}

type PongResponse struct {
	FromPeerId     PeerId
	FromOutPutAddr string
}

func NewPongResponse(fromPeerId PeerId, fromOutPutAddr string) *PongResponse {
	return &PongResponse{FromPeerId: fromPeerId, FromOutPutAddr: fromOutPutAddr}
}
