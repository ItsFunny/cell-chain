package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/enums"
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
	peerManager types.IPeerManager, bus eventbus.ICommonEventBus, discovery types.DiscoveryComponent) component.EnvelopeHandler {
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
	selfNode := p.peerManager.GetSelfNode()
	// TODO optimize:producer and consumer
	if p.peerManager.Have(from) {
		// send response
		remote := p.peerManager.GetByPeerId(from)
		resp := types.CreatePongEnvelopeResponse(p.cdc.GetCodec(), env.Header.SequenceId, selfNode.PeerId(), selfNode.MetaData().GetOutPutAddress())
		types.PublishDiscoverySendMessageEvent(p.bus, types.NewSendToPeerRequest(remote.MetaData().GetOutPutAddress(), resp))
	} else {
		// probe
		probeReq := types.CreateProbeEnvelopRequest(p.cdc.GetCodec(), env.Header.SequenceId, types.NewProbeRequest(selfNode.PeerId(), selfNode.MetaData()))
		types.PublishDiscoverySendMessageEvent(p.bus, types.NewSendToPeerRequest(from.ToString(), probeReq))
	}

	return nil
}

func (p *PingHandler) Protocol() string {
	return types.Ping
}
