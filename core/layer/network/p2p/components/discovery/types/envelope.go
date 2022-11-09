package types

import (
	"github.com/itsfunny/go-cell/base/common/utils"
	"github.com/itsfunny/go-cell/component/codec/types"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
)

func CreatePingPongEnvelopeRequest(cdc types.Codec) *types2.Envelope {
	seq := utils.GenerateSequenceId()
	req := PingPongRequest{}
	data, _ := cdc.Marshal(req)
	return types2.CreateNoopSignEnvelope(PingPong, seq, data)
}
