package typcore

import (
	"github.com/urfave/cli/v2"
)

// AppContext is context of app
type AppContext struct {
	*Descriptor
}

// ActionFunc to return ActionFunc to invoke function fn
func (a *AppContext) ActionFunc(fn interface{}) func(*cli.Context) error {
	return func(c *cli.Context) (err error) {
		return a.App.Invoke(a, c, fn)
	}
}
