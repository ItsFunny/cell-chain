package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/enums"
	types3 "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"github.com/itsfunny/go-cell/component/codec"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ component.EnvelopeHandler = (*PingHandler)(nil)
)

type PingHandler struct {
	ddd         *component.DDDComponent
	cdc         *codec.CodecComponent
	peerManager types.IPeerManager
	bus         eventbus.ICommonEventBus
}

func NewPingHandler(ddd *component.DDDComponent, cdc *codec.CodecComponent,
	peerManager types.IPeerManager, bus eventbus.ICommonEventBus, discovery types.DiscoveryComponent) *PingHandler {
	ret := &PingHandler{ddd: ddd, cdc: cdc, peerManager: peerManager, bus: bus}
	return ret
}

func (p *PingHandler) Handler(ctx *pipeline.Context, env *types2.Envelope) error {
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
		fromNode := p.peerManager.GetByPeerId(from)
		selfNode := p.peerManager.GetSelfNode()
		cellCtx := types3.CellContext{}
		cellCtx = cellCtx.FromPipelineCtx(ctx)
		//p.discovery.SendToPeerAsync(cellCtx, types.SendToPeerRequest{
		//	To:      fromNode,
		//	Envelop: types.CreatePongEnvelopeResponse(p.cdc.GetCodec(), env.Header.SequenceId, selfNode.PeerId(), selfNode.MetaData().GetOutPutAddress()),
		//})

	} else {
		// probe
	}
	return nil
}

func (p *PingHandler) Protocol() string {
	return types.Ping
}
