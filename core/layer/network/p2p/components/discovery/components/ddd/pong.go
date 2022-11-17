package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/component/codec"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ component.EnvelopeHandler = (*PongHandler)(nil)
)

type PongHandler struct {
	cdc         *codec.CodecComponent
	peerManager types.IPeerManager
}

func (p *PongHandler) Handler(ctx *pipeline.Context, env *types2.Envelope) error {
	data := env.Payload.Data
	var resp types.PongResponse
	if err := p.cdc.UnMarshal(data, &resp); nil != err {
		return err
	}
	// TODO, optimize
	p.peerManager.ForwardMessage() <- resp
	return nil
}

func (p *PongHandler) Protocol() string {
	return types.Pong
}
