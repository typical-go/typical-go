package typapp

import (
	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// Context of App
type Context struct {
	*typcore.AppContext
	*App
}

// ActionFunc to return ActionFunc to invoke function fn
func (c *Context) ActionFunc(fn interface{}) func(*cli.Context) error {
	return func(cliCtx *cli.Context) (err error) {
		return c.Invoke(cliCtx, fn)
	}
}

// Invoke function with Dependency Injection
func (c *Context) Invoke(cliCtx *cli.Context, fn interface{}) (err error) {
	di := dig.New()
	if cliCtx != nil {
		if err = di.Provide(func() *cli.Context { return cliCtx }); err != nil {
			return
		}
	}
	if c.Configuration != nil {
		// provide configuration to dependency-injection container
		if err = provide(di, c.Configuration.Provide()...); err != nil {
			return
		}
	}
	// provide registered function in descriptor to dependency-injection container
	if err = provide(di, c.Provide()...); err != nil {
		return
	}
	// invoke preparation as register in descriptor
	if err = invoke(di, c.Prepare()...); err != nil {
		return
	}

	startFn := func() error {
		return di.Invoke(fn)
	}
	stopFn := func() error {
		return invoke(di, c.Destroy()...)
	}
	for _, err := range common.StartGracefully(startFn, stopFn) {
		log.Error(err.Error())
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
