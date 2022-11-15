package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/component/codec"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ component.EnvelopeHandler = (*ProbeSyncHandler)(nil)
)

type ProbeSyncHandler struct {
	cdc         *codec.CodecComponent
	peerManager types.IPeerManager
}

func (p *ProbeSyncHandler) Handler(ctx *pipeline.Context, env *types2.Envelope) error {
	data := env.Payload.Data
	var resp types.ProbeResponse
	if err := p.cdc.UnMarshal(data, &resp); nil != err {
		return err
	}
	p.peerManager.Register(types.NewPeerWrapper(resp.SelfPeerId, resp.MetaData))
	return nil
}

func (p *ProbeSyncHandler) Protocol() string {
	return types.APIProbeResponse
}
