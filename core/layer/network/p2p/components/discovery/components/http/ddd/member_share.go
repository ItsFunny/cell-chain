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

func (m *MemberShareHandler) Handler(ctx *pipeline.Context, env *types2.Envelope) error {
	data := env.Payload.Data
	if len(data) == 0 {
		return nil
	}
	req := types.MembersShareRequest{}
	if err := m.cdc.UnMarshal(data, &req); nil != err {
		return err
	}

	remoteKnownPeers := req.KnownPeers
	knwonPeers := m.peerManager.GetMembership()
	logrusplugin.MInfo(enums.MemberShareHandler, "receive memshare message",
		"sequenceId", env.Header.SequenceId,
		"known peers", knwonPeers, "msg", req.String())

	unknownPeers := make(map[types.PeerId]string)
	for id, v := range remoteKnownPeers {
		_, exist := knwonPeers[id]
		if !exist {
			unknownPeers[id] = v
			continue
		}
	}

	removeUnknownPeers := make(map[types.PeerId]string)
	for id, v := range knwonPeers {
		_, exist := remoteKnownPeers[id]
		if !exist {
			removeUnknownPeers[id] = v.MetaData().GetOutPutAddress()
		}
	}
	// TODO sync metas
	return nil
}

func (m *MemberShareHandler) Protocol() string {
	return types.MembersShare
}
