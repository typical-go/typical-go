package typcli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/typical-go/typical-go/pkg/typcfg"
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
			// NOTE: Make sure the application is exit
			os.Exit(0)
		}()
		if c.ConfigLoader != nil {
			if err = provide(di, loaderFn(c.Context)); err != nil {
				return
			}
		}
		if configurer, ok := c.Module.(typcfg.Configurer); ok {
			_, _, loadFn := configurer.Configure()
			if err = provide(di, loadFn); err != nil {
				return
			}
		}
		return di.Invoke(fn)
	}
}
