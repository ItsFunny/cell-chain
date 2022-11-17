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
	_ component.EnvelopeHandler = (*MemberShareHandler)(nil)
)

type MemberShareHandler struct {
	cdc         *codec.CodecComponent
	bus         eventbus.ICommonEventBus
	peerManager types.IPeerManager
}

func NewMemberShareHandler(cdc *codec.CodecComponent, bus eventbus.ICommonEventBus, peerManager types.IPeerManager) component.EnvelopeHandler {
	return &MemberShareHandler{cdc: cdc, bus: bus, peerManager: peerManager}
}

func (m *MemberShareHandler) Handler(ctx *pipeline.Context, env *types2.Envelope) error {
	data := env.Payload.Data
	if len(data) == 0 {
		return nil
	}
	req := types.MembersShareRequest{}
	if err := m.cdc.UnMarshal(data, &req); nil != err {
		return err
	}

	seq := env.Header.SequenceId

	selfNode := m.peerManager.GetSelfNode()
	remoteKnownPeers := req.KnownPeers
	knwonPeers := m.peerManager.GetMembership()
	logrusplugin.MInfo(enums.MemberShareHandler, "receive memshare message",
		"sequenceId", env.Header.SequenceId,
		"known peers", knwonPeers, "msg", req.String())

	unknownPeers := make(map[types.PeerId]string)
	differPeers := make(map[types.PeerId]types.IPeerNode)
	for id, v := range remoteKnownPeers {
		node, exist := knwonPeers[id]
		if !exist {
			unknownPeers[id] = v
			continue
		}
		if v != node.MetaData().GetOutPutAddress() {
			differPeers[id] = node
		}
	}

	remoteUnknownPeers := make(map[types.PeerId]string)
	for id, v := range knwonPeers {
		_, exist := remoteKnownPeers[id]
		if !exist {
			remoteUnknownPeers[id] = v.MetaData().GetOutPutAddress()
		}
	}
	if len(unknownPeers) > 0 {
		// probe
		for _, v := range unknownPeers {
			proReq := types.NewProbeRequest(selfNode.PeerId(), selfNode.MetaData())
			types.PublishDiscoverySendMessageEvent(m.bus,
				types.NewSendToPeerRequest(v,
					types.CreateProbeEnvelopRequest(m.cdc.GetCodec(), seq, proReq)))
		}
	}

	if len(remoteUnknownPeers) > 0 {
		// help the remote node sync members
		for _, v := range remoteUnknownPeers {
			newMember := types.NewNewMemberRequest(v)
			types.PublishDiscoverySendMessageEvent(m.bus,
				types.NewSendToPeerRequest(v,
					types.CreateNewMemberEnvelopeRequest(m.cdc.GetCodec(), seq, newMember)))
		}
	}

	return nil
}

func (m *MemberShareHandler) Protocol() string {
	return types.MembersShare
}
