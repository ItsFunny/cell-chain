package component

import (
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/codec"
	"github.com/itsfunny/go-cell/component/codec/types"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	_ CellComponent = (*BaseComponent)(nil)
)

type CellMsgFn func() sdk.CellRequest

type CellComponent interface {
	services.IBaseService
	// TODO
	//Handlers() []EnvelopeHandler
}

type BaseComponent struct {
	*services.BaseService

	ddd *DDDComponent
	cdc *codec.CodecComponent
}

func NewBaseComponent(m logsdk.Module, impl CellComponent, ddd *DDDComponent, cdc *codec.CodecComponent) *BaseComponent {
	ret := &BaseComponent{}
	ret.BaseService = services.NewBaseService(nil, m, impl)
	ret.ddd = ddd
	ret.cdc = cdc
	return ret
}

func (c *BaseComponent) Send(ctx sdk.CellContext, fn CellMsgFn) (interface{}, error) {
	msgReq := fn()
	promise := c.ddd.Send(ctx, msgReq)
	return promise.GetForever()
}

func (c *BaseComponent) SendAsync(ctx sdk.CellContext, fn CellMsgFn) (*promise.Promise, error) {
	msgReq := fn()
	return c.ddd.Send(ctx, msgReq), nil
}

func (c *BaseComponent) GetCodec() types.Codec {
	return c.cdc.GetCodec()
}
