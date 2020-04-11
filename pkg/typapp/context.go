package typapp

import (
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// Context of App
type Context struct {
	*typcore.Descriptor
	*App

	container *dig.Container
}

// ActionFunc to return ActionFunc to invoke function fn
func (c *Context) ActionFunc(v interface{}) func(*cli.Context) error {
	return func(cliCtx *cli.Context) (err error) {
		if invokable, ok := v.(*Invocation); ok {
			return c.Invoke(cliCtx, invokable)
		}

		return c.Invoke(cliCtx, NewInvocation(v))
	}
}

// Invoke function with Dependency Injection
func (c *Context) Invoke(cliCtx *cli.Context, invocation *Invocation) (err error) {
	di := c.Container()

	if err = NewConstructor(func() *cli.Context {
		return cliCtx
	}).Provide(di); err != nil {
		return
	}

	for _, constructor := range c.Constructors() {
		if err = constructor.Provide(di); err != nil {
			return
		}
	}

	// invoke preparation as register in descriptor
	for _, preparation := range c.Preparations() {
		if err = preparation.Invoke(di); err != nil {
			return
		}
	}

	startFn := func() error { return invocation.Invoke(di) }

	for _, err := range common.StartGracefully(startFn, c.stop) {
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
	for _, destruction := range c.Destructions() {
		if err = destruction.Invoke(c.Container()); err != nil {
			return
		}
	}
	return
}
