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
	_ component.EnvelopeHandler = (*ProbeEnvelopeHandler)(nil)
)

type ProbeEnvelopeHandler struct {
	cdc         *codec.CodecComponent
	peerManager types.IPeerManager
	bus         eventbus.ICommonEventBus
}

func NewProbeEnvelopeHandler(cdc *codec.CodecComponent, peerManager types.IPeerManager, bus eventbus.ICommonEventBus) component.EnvelopeHandler {
	return &ProbeEnvelopeHandler{cdc: cdc, peerManager: peerManager, bus: bus}
}

func (p *ProbeEnvelopeHandler) Handler(ctx *pipeline.Context, env *types2.Envelope) error {
	data := env.Payload.Data
	var probeReq types.ProbeRequest
	if err := p.cdc.UnMarshal(data, &probeReq); nil != err {
		return err
	}
	if !p.peerManager.Have(probeReq.PeerId) {
		logrusplugin.MInfo(enums.ProbeHandler, "unknown probe", "msg", probeReq)
	}
	// send response to sync data
	selfNode := p.peerManager.GetSelfNode()
	probeResponse := types.NewProbeResponse(selfNode.PeerId(), selfNode.MetaData())
	types.PublishDiscoverySendMessageEvent(p.bus, types.NewSendToPeerRequest(probeReq.SelfMetaData.GetOutPutAddress(),
		types.CreateProbeEnvelopResponse(p.cdc.GetCodec(),
			env.Header.SequenceId, probeResponse)))
	return nil
}

func (p *ProbeEnvelopeHandler) Protocol() string {
	return types.Probe
}
