package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/component/codec"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ component.EnvelopeHandler = (*PingPongHandler)(nil)
)

type PingPongHandler struct {
	ddd *component.DDDComponent
	cdc *codec.CodecComponent
}

func (p *PingPongHandler) Handler(ctx *pipeline.Context, env *types2.Envelope) error {
	data := env.Payload.Data
	if len(data) == 0 {
		return nil
	}
	req := types.PingRequest{}
	if err := p.cdc.UnMarshal(data, &req); nil != err {
		return err
	}

	return nil
}

func (p *PingPongHandler) Protocol() string {
	return types.PingPong
}
