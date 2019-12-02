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

// ModuleCli command line module
type ModuleCli struct {
	*typctx.Context
}

// Action to return action function
func (c ModuleCli) Action(fn interface{}) func(ctx *cli.Context) error {
	return func(ctx *cli.Context) (err error) {
		di := dig.New()
		gracefulStop := make(chan os.Signal)
		signal.Notify(gracefulStop, syscall.SIGTERM)
		signal.Notify(gracefulStop, syscall.SIGINT)
		defer func() {
			gracefulStop <- syscall.SIGTERM
		}()
		if c.ConfigLoader != nil {
			if err = provide(di, func() typcfg.Loader { return c.ConfigLoader }); err != nil {
				return
			}
		}
		return di.Invoke(fn)
	}
}
