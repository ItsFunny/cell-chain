package types

import "context"

type CellContext struct {
	ctx context.Context
}

func (c CellContext) GetGoCtx() context.Context {
	return c.ctx
}

type CellRequest interface {
}

type CellResponse interface {
}
