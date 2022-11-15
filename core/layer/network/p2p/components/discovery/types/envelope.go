package types

import (
	"github.com/itsfunny/go-cell/base/common/utils"
	"github.com/itsfunny/go-cell/component/codec/types"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
)

func CreatePingEnvelopeRequest(cdc types.Codec, fromPeerId PeerId, fromOutPutAddr string) *types2.Envelope {
	seq := utils.GenerateSequenceId()
	req := NewPingRequest(fromPeerId, fromOutPutAddr)
	data, _ := cdc.Marshal(req)
	return types2.CreateNoopSignEnvelope(Ping, seq, data)
}

func CreatePongEnvelopeResponse(cdc types.Codec, seq string, fromPeerId PeerId, fromOutPutAddr string) *types2.Envelope {
	req := NewPongResponse(fromPeerId, fromOutPutAddr)
	data, _ := cdc.Marshal(req)
	return types2.CreateNoopSignEnvelope(Pong, seq, data)
}

func CreateMemberShareEnvelopeRequest(cdc types.Codec, fromPeerId PeerId, mems map[PeerId]IPeerNode) *types2.Envelope {
	seq := utils.GenerateSequenceId()
	memMetas := make(map[PeerId]string)
	for id, v := range mems {
		memMetas[id] = v.MetaData().GetOutPutAddress()
	}
	req := NewMembersShareRequest(fromPeerId, memMetas)
	data, _ := cdc.Marshal(req)
	return types2.CreateNoopSignEnvelope(MembersShare, seq, data)
}

func CreateProbeEnvelopRequest(cdc types.Codec, seq string, req *ProbeRequest) *types2.Envelope {
	data, _ := cdc.Marshal(req)
	return types2.CreateNoopSignEnvelope(Probe, seq, data)
}

func CreateProbeEnvelopResponse(cdc types.Codec, seq string, resp *ProbeResponse) *types2.Envelope {
	data, _ := cdc.Marshal(resp)
	return types2.CreateNoopSignEnvelope(APIProbeResponse, seq, data)
}
