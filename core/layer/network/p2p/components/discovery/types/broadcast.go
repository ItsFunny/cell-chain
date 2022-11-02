package types

import "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"

type BroadCastRequest struct {
	Envelop types.Envelope
}

type BroadCastResponse struct {
}

type SendToPeerRequest struct {
	To      IPeerNode
	Envelop types.Envelope
}
type SendToPeerResponse struct {
}
