package types

import (
	"context"
	"github.com/itsfunny/go-cell/base/common/utils"
	"github.com/itsfunny/go-cell/base/reactor"
	"github.com/itsfunny/go-cell/sdk/pipeline"
)

type HandlerFlag int

const (
	HandlerFlagNotify HandlerFlag = 1 << 0
	HandlerFlagRecord HandlerFlag = 1 << 1
)

type CellContext struct {
	ctx        context.Context
	SequenceId string
}

func (c CellContext) GetGoCtx() context.Context {
	return c.ctx
}

func (c CellContext) FromHttpCtx(ctx reactor.IBuzzContext) CellContext {
	c.SequenceId = ctx.GetCommandContext().Summary.GetSequenceId()
	c.ctx = ctx.GetCommandContext().ServerResponse.GetGoContext()
	return c
}

func (c CellContext) FromPipelineCtx(ctx *pipeline.Context) CellContext {
	return c
}

func EmptyCellContext(ctx context.Context) CellContext {
	ret := CellContext{
		ctx:        ctx,
		SequenceId: utils.GenerateSequenceId(),
	}
	return ret
}
func (c CellContext) WithSeq(s string) CellContext {
	c.SequenceId = s
	return c
}

type CellRequest interface {
}

type CellResponse interface {
}
