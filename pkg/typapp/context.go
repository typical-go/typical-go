package typapp

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

// Context of App
type Context struct {
	*typcore.Descriptor
	*TypicalApp

	container *typdep.Container
}

// ActionFunc to return ActionFunc to invoke function fn
func (c *Context) ActionFunc(v interface{}) func(*cli.Context) error {
	return func(cliCtx *cli.Context) (err error) {
		if invocation, ok := v.(*typdep.Invocation); ok {
			return c.Invoke(cliCtx, invocation)
		}
		return c.Invoke(cliCtx, typdep.NewInvocation(v))
	}
}

// Invoke function with Dependency Injection
func (c *Context) Invoke(cliCtx *cli.Context, invocation *typdep.Invocation) (err error) {
	di := c.Container()

	if err = typdep.Provide(di,
		typdep.NewConstructor(func() typcore.ConfigManager {
			return c.ConfigManager
		}),
		typdep.NewConstructor(func() *cli.Context {
			return cliCtx
		}),
	); err != nil {
		return
	}

	// provide registered function in descriptor to dependency-injection container
	if err = typdep.Provide(di, c.Provide()...); err != nil {
		return
	}

	// invoke preparation as register in descriptor
	if err = typdep.Invoke(di, c.Prepare()...); err != nil {
		return
	}

	startFn := func() error { return invocation.Invoke(di) }

	for _, err := range common.StartGracefully(startFn, c.stop) {
		log.Error(err.Error())
	}
	return
}

// Container for dependency-injection
func (c *Context) Container() *typdep.Container {
	if c.container == nil {
		typdep.New()
	}
	return c.container
}

func (c *Context) stop() (err error) {
	for _, destruction := range c.Destroy() {
		if err = destruction.Invocation.Invoke(c.Container()); err != nil {
			return
		}
	}
	return
}
