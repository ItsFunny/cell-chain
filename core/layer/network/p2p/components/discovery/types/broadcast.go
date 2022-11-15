package types

import "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"

type BroadCastRequest struct {
	Envelop *types.Envelope
}

type BroadCastResponse struct {
}

type SendToPeerRequest struct {
	To      string
	Envelop *types.Envelope
}

func NewSendToPeerRequest(to string, envelop *types.Envelope) *SendToPeerRequest {
	return &SendToPeerRequest{To: to, Envelop: envelop}
}

type SendToPeerResponse struct {
}
