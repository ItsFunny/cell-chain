package component

import (
	"github.com/itsfunny/cell-chain/common/enums"
	"github.com/itsfunny/cell-chain/common/types"
	"github.com/itsfunny/go-cell/base/core/promise"
	"github.com/itsfunny/go-cell/base/core/services"
	"github.com/itsfunny/go-cell/component/routine"
	"github.com/itsfunny/go-cell/sdk/pipeline"
	"reflect"
)

type DDDHandler interface {
	Handler(ctx *pipeline.Context) (types.CellResponse, error)
	PredictMsg() types.CellRequest
}

type DDDComponent struct {
	*services.BaseService
	pip     pipeline.Pipeline
	routine routine.IRoutineComponent
}

func NewDDDComponent(routine routine.IRoutineComponent) *DDDComponent {
	ret := &DDDComponent{}
	ret.BaseService = services.NewBaseService(nil, enums.DDDModule, ret)
	ret.pip = pipeline.New()
	ret.routine = routine
	return ret
}

func (d *DDDComponent) OnStart(ctx *services.StartCTX) error {
	return nil
}

func (d *DDDComponent) RegisterDDDHandler(h DDDHandler) {
	msg := h.PredictMsg()
	d.pip.RegisterFunc(reflect.TypeOf(msg), func(ctx *pipeline.Context) {
		d.routine.AddJob(func() {
			ret, err := h.Handler(ctx)
			if nil != err {
				ctx.Promise.Fail(err)
			} else {
				ctx.Promise.TrySend(ret)
			}
		})
	})
}

func (d *DDDComponent) Send(ctx types.CellContext, msg types.CellRequest) *promise.Promise {
	return d.pip.Serve(ctx.GetGoCtx(), msg)
}
