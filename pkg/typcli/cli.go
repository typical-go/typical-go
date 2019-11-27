package typcli

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typmodule"
	"github.com/urfave/cli"
	"go.uber.org/dig"
)

// Cli interface
type Cli interface {
	Action(fn interface{}) func(ctx *cli.Context) error
}

// NewCli return new constructor
func NewCli(ctx *typctx.Context, obj interface{}) Cli {
	return &cliImpl{
		Context: ctx,
		Obj:     obj,
	}
}

type cliImpl struct {
	*typctx.Context
	Obj interface{}
}

// Action to return action function
func (c cliImpl) Action(fn interface{}) func(ctx *cli.Context) error {
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
				fmt.Println("Error: " + err.Error())
				os.Exit(1)
				return
			}
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

func (c cliImpl) provideDependency(di *dig.Container) (err error) {
	if c.ConfigLoader != nil {
		if err = provide(di, func() typcfg.Loader { return c.ConfigLoader }); err != nil {
			return
		}
	}
	if provider, ok := c.Obj.(typmodule.Provider); ok {
		if err = provide(di, provider.Provide()...); err != nil {
			return
		}
	}
	return
}

func (c cliImpl) prepare(di *dig.Container) (err error) {
	if preparer, ok := c.Obj.(typmodule.Preparer); ok {
		if err = invoke(di, preparer.Prepare()...); err != nil {
			return
		}
	}
	return
}

func (c cliImpl) shutdown(di *dig.Container) (err error) {
	if destroyer, ok := c.Obj.(typmodule.Destroyer); ok {
		if err = invoke(di, destroyer.Destroy()...); err != nil {
			return
		}
	}
	return
}
