package types

import (
	"context"
	"github.com/itsfunny/go-cell/base/reactor"
	"gitlab.ebidsun.com/chain/droplib/plugin/pipeline"
)

type HandlerFlag int

const (
	HandlerFlagNotify HandlerFlag = 1 << 0
	HandlerFlagRecord HandlerFlag = 1 << 1
)

type CellContext struct {
	ctx context.Context
}

func (c CellContext) GetGoCtx() context.Context {
	return c.ctx
}

func (c CellContext) FromHttpCtx(ctx reactor.IBuzzContext) CellContext {
	return c
}

func (c CellContext) FromPipelineCtx(ctx *pipeline.Context) CellContext {
	return c
}

func EmptyCellContext(ctx context.Context) CellContext {
	ret := CellContext{
		ctx: ctx,
	}
	return ret
}

type CellRequest interface {
}

type CellResponse interface {
}
