package execkit

import "context"

type (
	// Runner responsble to run
	Runner interface {
		Run(context.Context) error
	}
	// RunFn run function
	RunFn      func(context.Context) error
	runnerImpl struct {
		fn RunFn
	}
)

// Run the runner
func Run(ctx context.Context, runner Runner) error {
	return runner.Run(ctx)
}

//
// runnerImpl
//

// NewRunner return new instance of runners
func NewRunner(fn RunFn) Runner {
	return &runnerImpl{fn: fn}
}

func (r *runnerImpl) Run(ctx context.Context) error {
	return r.fn(ctx)
}
