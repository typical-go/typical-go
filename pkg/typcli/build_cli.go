package typcli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// BuildCli command line module
type BuildCli struct {
	*typctx.Context
	Module interface{}
}

// Action to return action function
func (c BuildCli) Action(fn interface{}) func(ctx *cli.Context) error {
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
			os.Exit(0) // NOTE: Make sure the application is exit
		}()
		if err = provideLoader(di, c.Context); err != nil {
			return
		}
		if err = provideConfigFn(di, c.Module); err != nil {
			return
		}
		return di.Invoke(fn)
	}
}

// PreparedAction to return function with preparation, provide and destroy dependencies from other module
func (c BuildCli) PreparedAction(fn interface{}) func(ctx *cli.Context) error {
	return preparedAction(fn, c.Context)
}
