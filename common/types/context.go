package types

import (
	"context"
	"github.com/itsfunny/go-cell/base/reactor"
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

type CellRequest interface {
}

type CellResponse interface {
}
