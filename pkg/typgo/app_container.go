package typgo

import (
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// AppContainer of App
type AppContainer struct {
	di *dig.Container
}

// ActionFunc to return ActionFunc to invoke function fn
func (c *AppContainer) ActionFunc(v interface{}) func(*cli.Context) error {
	return func(cliCtx *cli.Context) (err error) {
		return c.Invoke(cliCtx, v)
	}
}

// Invoke function with Dependency Injection
func (c *AppContainer) Invoke(cliCtx *cli.Context, fn interface{}) (err error) {

	ctor := &Constructor{
		Fn: func() *cli.Context {
			return cliCtx
		},
	}

	if err = provide(c.di, ctor); err != nil {
		return
	}

	for _, ctor := range _ctors {
		if err = provide(c.di, ctor); err != nil {
			return
		}
	}

	startFn := func() error { return c.di.Invoke(fn) }

	common.StartGracefuly(startFn, c.stop)
	// for _, err := range common.StartGracefuly(startFn, c.stop) {
	// c.Warn(err.Error())
	// }
	return
}

func (c *AppContainer) stop() (err error) {
	for _, dtor := range _dtors {
		if err = c.di.Invoke(dtor.Fn); err != nil {
			return
		}
	}
	return
}

func provide(di *dig.Container, c *Constructor) (err error) {
	if c.Fn == nil {
		panic("provide: Fn is missing")
	}
	return di.Provide(c.Fn, dig.Name(c.Name))
}
