package types

import (
	"github.com/itsfunny/go-cell/base/common/utils"
	"github.com/itsfunny/go-cell/component/codec/types"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
)

func CreatePingPongEnvelopeRequest(cdc types.Codec, fromPeerId PeerId, fromOutPutAddr string) *types2.Envelope {
	seq := utils.GenerateSequenceId()
	req := NewPingRequest(fromPeerId, fromOutPutAddr)
	data, _ := cdc.Marshal(req)
	return types2.CreateNoopSignEnvelope(PingPong, seq, data)
}
