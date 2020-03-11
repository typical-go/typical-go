package exor

import (
	"context"
)

// Executor responsible to execution
type Executor interface {
	Execute(context.Context) error
}

// New executor
func New(fn func(context.Context) error) Executor {
	return &executor{fn: fn}
}

// Execute all executor
func Execute(ctx context.Context, executors ...Executor) (err error) {
	for _, executor := range executors {
		if err = executor.Execute(ctx); err != nil {
			return
		}
	}
	return
}

type executor struct {
	fn func(context.Context) error
}

func (e *executor) Execute(ctx context.Context) error {
	return e.fn(ctx)
}
