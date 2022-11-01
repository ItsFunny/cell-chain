package component

import (
	sdk "github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/routine"
	logsdk "github.com/itsfunny/go-cell/sdk/log"
)

var (
	_ CellComponent = (*BaseComponent)(nil)
)

type CellMsgFn func(msg interface{}) sdk.CellMsg

type CellComponent interface {
	services.IBaseService
}

type BaseComponent struct {
	*services.BaseService

	routinePool routine.IRoutineComponent
}

func NewBaseComponent(m logsdk.Module, impl CellComponent) *BaseComponent {
	ret := &BaseComponent{}
	ret.BaseService = services.NewBaseService(nil, m, impl)
	return ret
}

func (c *BaseComponent) Send(ctx sdk.CellContext, fn CellMsgFn) {
	ch := make(chan interface{})
	c.routinePool.AddJob(func() {

	})
}
