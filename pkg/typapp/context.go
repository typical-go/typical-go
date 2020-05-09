package typapp

import (
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// Context of App
type Context struct {
	*typcore.Descriptor
	typlog.Logger

	App       *App
	container *dig.Container
}

// ActionFunc to return ActionFunc to invoke function fn
func (c *Context) ActionFunc(v interface{}) func(*cli.Context) error {
	return func(cliCtx *cli.Context) (err error) {
		return c.Invoke(cliCtx, v)
	}
}

// Invoke function with Dependency Injection
func (c *Context) Invoke(cliCtx *cli.Context, fn interface{}) (err error) {
	di := c.Container()

	ctor := &Constructor{
		Fn: func() *cli.Context {
			return cliCtx
		},
	}

	if err = provide(di, ctor); err != nil {
		return
	}

	for _, ctor := range c.App.Constructors() {
		if err = provide(di, ctor); err != nil {
			return
		}
	}

	startFn := func() error { return di.Invoke(fn) }

	for _, err := range common.StartGracefuly(startFn, c.stop) {
		c.Warn(err.Error())
	}
	return
}

// Container for dependency-injection
func (c *Context) Container() *dig.Container {
	if c.container == nil {
		c.container = dig.New()
	}
	return c.container
}

func (c *Context) stop() (err error) {
	for _, destructor := range c.App.Destructors() {
		if err = destructor.Invoke(c.Container()); err != nil {
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
