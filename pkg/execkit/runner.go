package execkit

import "context"

// Runner responsble to run
type Runner interface {
	Run(context.Context) error
}
