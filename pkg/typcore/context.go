package typcore

import (
	log "github.com/sirupsen/logrus"

	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// NewContext return new instance of Context
func NewContext(d *ProjectDescriptor, obj interface{}) *Context {
	return &Context{
		ProjectDescriptor: d,
		obj:               obj,
	}
}

// Context of application
type Context struct {
	*ProjectDescriptor
	obj interface{}
}

// Action to return action function that required config and object only
func (c *Context) Action(fn interface{}) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		if err = c.provideConfigLoader(di); err != nil {
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
func (c *Context) PreparedAction(fn interface{}) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		if err = c.provideConfigLoader(di); err != nil {
			return
		}
		if err = c.provideAll(di); err != nil {
			return
		}
		if err = c.prepareAll(di); err != nil {
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
		if err = c.destroyAll(di); err != nil {
			log.Error(err.Error())
		}
		return
	}
}

func (c *Context) provideConfigLoader(di *dig.Container) (err error) {
	fn := func() ConfigLoader {
		return c.ConfigLoader
	}
	if c.ConfigLoader != nil {
		if err = provide(di, fn); err != nil {
			return
		}
	}
	return
}

func (c *Context) destroyAll(di *dig.Container) (err error) {
	for _, module := range c.AllModule() {
		if destroyer, ok := module.(Destroyer); ok {
			if err = invoke(di, destroyer.Destroy()...); err != nil {
				return
			}
		}
	}
	return
}

func (c *Context) provideAll(di *dig.Container) (err error) {
	if err = provide(di, c.Constructors...); err != nil {
		return
	}
	for _, module := range c.AllModule() {
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

func (c *Context) prepareAll(di *dig.Container) (err error) {
	for _, module := range c.AllModule() {
		if preparer, ok := module.(Preparer); ok {
			if err = invoke(di, preparer.Prepare()...); err != nil {
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

func provideConfigFn(di *dig.Container, v interface{}) (err error) {
	if configurer, ok := v.(Configurer); ok {
		_, _, loadFn := configurer.Configure()
		if err = provide(di, loadFn); err != nil {
			return
		}
	}
	return
}
