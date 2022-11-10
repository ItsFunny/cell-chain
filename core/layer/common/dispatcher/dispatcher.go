package dispatcher

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/enums"
	"github.com/itsfunny/cell-chain/common/types"
	types4 "github.com/itsfunny/cell-chain/core/layer/common/types"
	logrusplugin "github.com/itsfunny/go-cell/sdk/log/logrus"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ component.DDDHandler = (*MsgDispatcher)(nil)
)

type MsgDispatcher struct {
	handlers map[string]component.EnvelopeHandler
}

// TODO, use fx#invoke
func NewMsgDispatcher() *MsgDispatcher {
	ret := &MsgDispatcher{
		handlers: make(map[string]component.EnvelopeHandler),
	}
	return ret
}

func (m *MsgDispatcher) Handler(ctx *pipeline.Context) (types.CellResponse, types.HandlerFlag, error) {
	req := ctx.Request.(*types4.Peer2PeerRequest)
	env := req.Envelope
	protocol := env.Header.Protocol
	h := m.handlers[protocol]
	if h == nil {
		logrusplugin.MWarn(enums.HTTPEnvelopeDispatcher, "unknown envelop", protocol)
		return nil, types.HandlerFlagNotify, nil
	}

	err := h.Handler(ctx, env)
	return nil, types.HandlerFlagNotify, err
}

func (m *MsgDispatcher) PredictMsg() types.CellRequest {
	return &types4.Peer2PeerRequest{}
}
