package component

import (
	"github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/sdk/pipeline"
	"reflect"
)

type DDDComponent struct {
	*services.BaseService
	pip pipeline.Pipeline
}

type DDDWrapper struct {
}

type DDDHandler func(ctx *pipeline.Context) (types.CellResponse, error)

func (d *DDDComponent) OnStart(ctx *services.StartCTX) error {
	return nil
}

func (d *DDDComponent) RegisterDDDHandler(msg reflect.Type, h DDDHandler) {
	d.pip.RegisterFunc(msg, func(ctx *pipeline.Context) {
		ret, err := h(ctx)
		if nil != err {
			ctx.Promise.Fail(err)
		} else {
			ctx.Promise.TrySend(ret)
		}
	})
}

func (d *DDDComponent) Send(ctx types.CellContext, msg types.CellRequest) *promise.Promise {
	return d.pip.Serve(ctx.GetGoCtx(), msg)
}
