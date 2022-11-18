package ddd

import (
	"bytes"
	"github.com/itsfunny/cell-chain/common/component"
	"github.com/itsfunny/cell-chain/common/types"
	types3 "github.com/itsfunny/cell-chain/core/layer/common/types"
	types2 "github.com/itsfunny/cell-chain/core/layer/network/p2p/components/discovery/components/http/types"
	"github.com/itsfunny/go-cell/component/codec"
	"github.com/itsfunny/go-cell/sdk/pipeline"
	"net/http"
)

var (
	_ component.DDDHandler = (*HttpSendPeerDDDHandler)(nil)
)

type HttpSendPeerDDDHandler struct {
	cdcComponent *codec.CodecComponent
}

func NewHttpSendPeerDDDHandler(cdcComponent *codec.CodecComponent) component.DDDHandler {
	return &HttpSendPeerDDDHandler{cdcComponent: cdcComponent}
}

func (h HttpSendPeerDDDHandler) Handler(ctx *pipeline.Context) (types.CellResponse, types.HandlerFlag, error) {
	req := ctx.Request.(types2.HttpSendToPeerRequest)
	msg := req.Envelop
	to := req.To
	p2pReq := types3.NewPeer2PeerRequest(msg)
	err := Send(p2pReq, h.cdcComponent, to)

	return nil, types.HandlerFlagNotify, err
}

func (h HttpSendPeerDDDHandler) PredictMsg() types.CellRequest {
	return types2.HttpSendToPeerRequest{}
}

func Send(p *types3.Peer2PeerRequest, cdc *codec.CodecComponent, address string) error {
	data := cdc.MustMarshal(p)
	// TODO ,handle response
	_, err := http.Post(address, "application/json;charset=utf-8", bytes.NewBuffer(data))
	return err
}
