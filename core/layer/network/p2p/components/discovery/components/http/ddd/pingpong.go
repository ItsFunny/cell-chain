package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/enums"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/component/codec"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ component.EnvelopeHandler = (*PingPongHandler)(nil)
)

type PingPongHandler struct {
	ddd         *component.DDDComponent
	cdc         *codec.CodecComponent
	peerManager types.IPeerManager
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
	from := req.FromPeerId
	logrusplugin.MInfo(enums.PingPongHandler, "receive message",
		"from", from, "fromOutPutAddr", req.FromOutPutAddr)
	if p.peerManager.Have(from) {
		// send response

	} else {
		// probe
	}
	return nil
}

func (p *PingPongHandler) Protocol() string {
	return types.PingPong
}
