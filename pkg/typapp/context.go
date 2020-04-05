package typapp

import (
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

// Context of App
type Context struct {
	*typcore.Descriptor
	*App

	container *typdep.Container
}

// ActionFunc to return ActionFunc to invoke function fn
func (c *Context) ActionFunc(v interface{}) func(*cli.Context) error {
	return func(cliCtx *cli.Context) (err error) {
		if invokable, ok := v.(typdep.Invokable); ok {
			return c.Invoke(cliCtx, invokable)
		}

		return c.Invoke(cliCtx, typdep.NewInvocation(v))
	}
}

// Invoke function with Dependency Injection
func (c *Context) Invoke(cliCtx *cli.Context, invokable typdep.Invokable) (err error) {
	di := c.Container()

	if err = typdep.Provide(di,
		typdep.NewConstructor(func() *cli.Context {
			return cliCtx
		}),
	); err != nil {
		return
	}

	for _, constructor := range c.Constructors() {
		if err = constructor.Constructor.Provide(di); err != nil {
			return
		}
	}

	// invoke preparation as register in descriptor
	for _, preparation := range c.Preparations() {
		if err = preparation.Invoke(di); err != nil {
			return
		}
	}

	startFn := func() error { return invokable.Invoke(di) }

	for _, err := range common.StartGracefully(startFn, c.stop) {
		c.Warn(err.Error())
	}
	return
}

// Container for dependency-injection
func (c *Context) Container() *typdep.Container {
	if c.container == nil {
		c.container = typdep.New()
	}
	return c.container
}

func (c *Context) stop() (err error) {
	for _, destruction := range c.Destructions() {
		if err = destruction.Invocation.Invoke(c.Container()); err != nil {
			return
		}
	}
	return
}
