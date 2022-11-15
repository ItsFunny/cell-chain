package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/types"
	"github.com/itsfunny/go-cell/base/core/eventbus"
	"github.com/itsfunny/go-cell/component/codec"
	types2 "github.com/itsfunny/go-cell/framework/rpc/grpc/common/types"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ component.EnvelopeHandler = (*NewMemberHandler)(nil)
)

type NewMemberHandler struct {
	cdc         *codec.CodecComponent
	peerManager types.IPeerManager
	bus         eventbus.ICommonEventBus
}

func (n *NewMemberHandler) Handler(ctx *pipeline.Context, env *types2.Envelope) error {
	data := env.Payload.Data
	var req types.NewMemberRequest
	if err := n.cdc.UnMarshal(data, &req); nil != err {
		return err
	}
	seq := env.Header.SequenceId
	// TODO, probe or just fill with the data
	selfNode := n.peerManager.GetSelfNode()
	probeReq := types.NewProbeRequest(selfNode.PeerId(), selfNode.MetaData())
	types.PublishDiscoverySendMessageEvent(n.bus,
		types.NewSendToPeerRequest(req.PeerOutPutAddress,
			types.CreateProbeEnvelopRequest(n.cdc.GetCodec(), seq, probeReq)))
	return nil
}

func (n *NewMemberHandler) Protocol() string {
	return types.APINewMember
}
