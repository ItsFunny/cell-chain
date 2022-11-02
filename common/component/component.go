package component

import (
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/core/services"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	_ CellComponent = (*BaseComponent)(nil)
)

type CellMsgFn func() sdk.CellRequest

type CellComponent interface {
	services.IBaseService
}

type BaseComponent struct {
	*services.BaseService

	ddd *DDDComponent
}

func NewBaseComponent(m logsdk.Module, impl CellComponent, ddd *DDDComponent) *BaseComponent {
	ret := &BaseComponent{}
	ret.BaseService = services.NewBaseService(nil, m, impl)
	ret.ddd = ddd
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
