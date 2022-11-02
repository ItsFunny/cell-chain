package types

import (
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
)

type HttpSendToPeerRequest struct {
	To      types.IPeerNode
	Envelop types2.Envelope
}

func (h HttpSendToPeerRequest) From(request types.SendToPeerRequest) HttpSendToPeerRequest {
	h.To = request.To
	h.Envelop = request.Envelop
	return h
}
