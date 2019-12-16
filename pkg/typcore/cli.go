package typcore

import (
	log "github.com/sirupsen/logrus"

	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// Cli for command line
type Cli interface {
	Action(fn interface{}) func(ctx *cli.Context) error
	PreparedAction(fn interface{}) func(ctx *cli.Context) error
	Context() *Context
	Object() interface{}
}

// NewCli return new instance of Cli
func NewCli(ctx *Context, obj interface{}) Cli {
	return &cliImpl{
		ctx: ctx,
		obj: obj,
	}
}

type cliImpl struct {
	ctx *Context
	obj interface{}
}

// Action to return action function that required config and object only
func (c *cliImpl) Action(fn interface{}) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		if err = provideLoader(di, c.ctx); err != nil {
			return
		}
		if err = provideConfigFn(di, c.obj); err != nil {
			return
		}
		gracefulStop := make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)
		go func() {
			defer func() {
				gracefulStop <- syscall.SIGTERM
			}()
			err = di.Invoke(fn)
		}()
		<-gracefulStop
		if err != nil {
			log.Error(err.Error())
		}
		return
	}
}

// PreparedAction to return function with preparation, provide and destroy dependencies from other module
func (c *cliImpl) PreparedAction(fn interface{}) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		if err = provideAll(di, c.ctx); err != nil {
			return
		}
		if err = prepareAll(di, c.ctx); err != nil {
			return
		}
		gracefulStop := make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)
		go func() {
			defer func() {
				gracefulStop <- syscall.SIGTERM
			}()
			err = di.Invoke(fn)
		}()
		<-gracefulStop
		if err != nil {
			log.Error(err.Error())
		}
		if err = destroyAll(di, c.ctx); err != nil {
			log.Error(err.Error())
		}
		return
	}
}

// Context of Cli
func (c *cliImpl) Context() *Context {
	return c.ctx
}

// Object of Cli
func (c *cliImpl) Object() interface{} {
	return c.obj
}

func provideAll(di *dig.Container, ctx *Context) (err error) {
	if err = provideLoader(di, ctx); err != nil {
		return
	}
	if err = provide(di, ctx.Constructors...); err != nil {
		return
	}
	for _, module := range ctx.AllModule() {
		if err = provideConfigFn(di, module); err != nil {
			return
		}
		if provider, ok := module.(Provider); ok {
			if err = provide(di, provider.Provide()...); err != nil {
				return
			}
		}
	}
	return
}

func prepareAll(di *dig.Container, ctx *Context) (err error) {
	for _, module := range ctx.AllModule() {
		if preparer, ok := module.(Preparer); ok {
			if err = invoke(di, preparer.Prepare()...); err != nil {
				return
			}
		}
	}
	return
}

func destroyAll(di *dig.Container, ctx *Context) (err error) {
	for _, module := range ctx.AllModule() {
		if destroyer, ok := module.(Destroyer); ok {
			if err = invoke(di, destroyer.Destroy()...); err != nil {
				return
			}
		}
	}
	return
}

func invoke(di *dig.Container, fns ...interface{}) (err error) {
	for _, fn := range fns {
		if err = di.Invoke(fn); err != nil {
			return
		}
	}
	return
}

func provide(di *dig.Container, fns ...interface{}) (err error) {
	for _, fn := range fns {
		if err = di.Provide(fn); err != nil {
			return
		}
	}
	return
}

func loaderFn(c *Context) interface{} {
	return func() ConfigLoader {
		return c.ConfigLoader
	}
}

func provideLoader(di *dig.Container, ctx *Context) (err error) {
	if ctx.ConfigLoader != nil {
		if err = provide(di, loaderFn(ctx)); err != nil {
			return
		}
	}
	return
}

func provideConfigFn(di *dig.Container, v interface{}) (err error) {
	if configurer, ok := v.(Configurer); ok {
		_, _, loadFn := configurer.Configure()
		if err = provide(di, loadFn); err != nil {
			return
		}
	}
	return
}
