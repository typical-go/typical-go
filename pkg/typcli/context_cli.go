package typcli

import (
	"os"
	"os/signal"
	"syscall"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typmodule"
	"github.com/urfave/cli"
	"go.uber.org/dig"
)

// NewContextCli return new instance of context cli
func NewContextCli(ctx *typctx.Context) Cli {
	return &contextCli{Context: ctx}
}

type contextCli struct {
	*typctx.Context
}

// Action to return action function
func (c *contextCli) Action(fn interface{}) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		gracefulStop := make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)
		defer func() {
			gracefulStop <- syscall.SIGTERM
		}()
		go func() {
			<-gracefulStop
			if err = c.shutdown(di); err != nil {
				log.Fatal(err.Error())
			}
			// NOTE: Make sure the application is exit
			os.Exit(0)
		}()
		if err = c.provideDependency(di); err != nil {
			return
		}
		if err = c.prepare(di); err != nil {
			return
		}
		return di.Invoke(fn)
	}
}

func (c *contextCli) provideDependency(di *dig.Container) (err error) {
	if c.ConfigLoader != nil {
		if err = provide(di, func() typcfg.Loader { return c.ConfigLoader }); err != nil {
			return
		}
	}
	if err = provide(di, c.Constructors...); err != nil {
		return
	}
	for _, module := range c.AllModule() {
		if provider, ok := module.(typmodule.Provider); ok {
			if err = provide(di, provider.Provide()...); err != nil {
				return
			}
		}
	}
	return
}

func (c *contextCli) prepare(di *dig.Container) (err error) {
	for _, module := range c.AllModule() {
		if preparer, ok := module.(typmodule.Preparer); ok {
			if err = invoke(di, preparer.Prepare()...); err != nil {
				return
			}
		}
	}
	return
}

func (c *contextCli) shutdown(di *dig.Container) (err error) {
	for _, module := range c.AllModule() {
		if destroyer, ok := module.(typmodule.Destroyer); ok {
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
