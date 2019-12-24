package typcore

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"

	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// NewContext return new instance of Context
func NewContext(d *ProjectDescriptor) *Context {
	return &Context{
		ProjectDescriptor: d,
	}
}

// Context of application
type Context struct {
	*ProjectDescriptor
}

// Action to return action function that required config and object only
func (c *Context) Action(obj, fn interface{}) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		if err = c.provideConfigLoader(di); err != nil {
			return
		}
		if err = provideLoadConfigFn(di, obj); err != nil {
			return
		}
		if err = provideCliContextFn(di, ctx); err != nil {
			return
		}
		app := common.Application{
			StartFn: func() error { return di.Invoke(fn) },
		}
		for _, err := range app.Run() {
			log.Error(err.Error())
		}
		return
	}
}

// PreparedAction to return function with preparation, provide and destroy dependencies from other module
func (c *Context) PreparedAction(fn interface{}) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		if err = provide(di, c.Constructors...); err != nil {
			return
		}
		if err = c.provideConfigLoader(di); err != nil {
			return
		}
		if err = c.provideModules(di); err != nil {
			return
		}
		if err = c.prepareModules(di); err != nil {
			return
		}
		if err = provideCliContextFn(di, ctx); err != nil {
			return
		}
		runner := common.Application{
			StartFn: func() error { return di.Invoke(fn) },
			StopFn:  func() error { return c.destroyModules(di) },
		}
		for _, err := range runner.Run() {
			log.Error(err.Error())
		}
		return
	}
}

func (c *Context) destroyModules(di *dig.Container) (err error) {
	for _, module := range c.AllModule() {
		if destroyer, ok := module.(Destroyer); ok {
			if err = invoke(di, destroyer.Destroy()...); err != nil {
				return
			}
		}
	}
	return
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

func (c *Context) provideModules(di *dig.Container) (err error) {
	for _, module := range c.AllModule() {
		if err = provideLoadConfigFn(di, module); err != nil {
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

func (c *Context) prepareModules(di *dig.Container) (err error) {
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

func provideLoadConfigFn(di *dig.Container, v interface{}) (err error) {
	if configurer, ok := v.(Configurer); ok {
		_, _, loadFn := configurer.Configure()
		if err = provide(di, loadFn); err != nil {
			return
		}
	}
	return
}

func provideCliContextFn(di *dig.Container, ctx *cli.Context) error {
	fn := func() *cli.Context {
		return ctx
	}
	return provide(di, fn)
}
