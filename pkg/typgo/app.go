package typgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

// App information
type App struct {
	di *dig.Container
}

// ActionFunc to return ActionFunc to invoke function fn
func (c *App) ActionFunc(v interface{}) func(*cli.Context) error {
	return func(cliCtx *cli.Context) (err error) {
		return c.Invoke(cliCtx, v)
	}
}

// Invoke function with Dependency Injection
func (c *App) Invoke(cliCtx *cli.Context, fn interface{}) (err error) {

	ctor := &Constructor{
		Fn: func() *cli.Context {
			return cliCtx
		},
	}

	if err = provide(c.di, ctor); err != nil {
		return
	}

	for _, ctor := range _ctors {
		if err = provide(c.di, ctor); err != nil {
			return
		}
	}

	startFn := func() error { return c.di.Invoke(fn) }

	common.StartGracefuly(startFn, c.stop)
	// for _, err := range common.StartGracefuly(startFn, c.stop) {
	// c.Warn(err.Error())
	// }
	return
}

func (c *App) stop() (err error) {
	for _, dtor := range _dtors {
		if err = c.di.Invoke(dtor.Fn); err != nil {
			return
		}
	}
	return
}

func provide(di *dig.Container, c *Constructor) (err error) {
	if c.Fn == nil {
		panic("provide: Fn is missing")
	}
	return di.Provide(c.Fn, dig.Name(c.Name))
}

func createApp(d *Descriptor) *App {
	di := dig.New()
	di.Provide(func() *Descriptor {
		return d
	})
	return &App{di: di}
}

func createAppCli(d *Descriptor) *cli.App {
	a := createApp(d)

	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Before = func(*cli.Context) (err error) {
		if configFile := os.Getenv("CONFIG"); configFile != "" {
			_, err = typcfg.Load(configFile)
		}
		return
	}
	app.Version = d.Version
	app.Action = a.ActionFunc(d.EntryPoint)
	return app
}
