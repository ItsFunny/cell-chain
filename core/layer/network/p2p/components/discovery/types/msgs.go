package types

type PingRequest struct {
	FromPeerId     PeerId
	FromOutPutAddr string
}

func NewPingRequest(fromPeerId PeerId, fromOutPutAddr string) *PingRequest {
	return &PingRequest{FromPeerId: fromPeerId, FromOutPutAddr: fromOutPutAddr}
}
