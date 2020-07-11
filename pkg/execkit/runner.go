package execkit

import "context"

// Runner responsble to run
type Runner interface {
	Run(context.Context) error
}

// Run the runner
func Run(ctx context.Context, runner Runner) error {
	return runner.Run(ctx)
}
