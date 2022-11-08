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
	to := req.To
	p2pReq := types2.NewPeer2PeerRequest(msg)
	err := p2pReq.Send(h.cdcComponent, to.MetaData().GetOutPutAddress())

	return nil, types.HandlerFlagNotify, err
}

func (h HttpSendPeerDDDHandler) PredictMsg() types.CellRequest {
	return types2.HttpSendToPeerRequest{}
}
