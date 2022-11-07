package ddd

import (
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/types"
	types2 "github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/types"
	"github.com/itsfunny/go-cell/component/codec"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

var (
	_ component.DDDHandler = (*HttpSendPeerDDDHandler)(nil)
)

type HttpSendPeerDDDHandler struct {
	cdcComponent *codec.CodecComponent
}

func (h HttpSendPeerDDDHandler) Handler(ctx *pipeline.Context) (types.CellResponse, types.HandlerFlag, error) {
	req := ctx.Request.(types2.HttpSendToPeerRequest)
	msg := req.Envelop
	marshal, err := h.cdcComponent.Marshal(msg)
	if nil != err {
		return nil, types.HandlerFlagNotify, err
	}
	to := req.To

	return nil, types.HandlerFlagNotify, nil
}

func (h HttpSendPeerDDDHandler) PredictMsg() types.CellRequest {
	return types2.HttpSendToPeerRequest{}
}
