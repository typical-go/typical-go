package typapp

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// App is the cli app
func (a *App) cli(d *typcore.Descriptor) *cli.App {
	c := &Context{
		Descriptor: d,
		App:        a,
	}
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Before = func(*cli.Context) (err error) {
		if configFile := os.Getenv("CONFIG"); configFile != "" {
			c.Info("Load environments")
			_, err = typcfg.Load(configFile)
		}
		return
	}
	app.Version = d.Version
	app.Action = c.ActionFunc(a.main)
	app.Commands = a.Commands(c)
	return app
}

// RunApp to run the applciation
func (a *App) RunApp(d *typcore.Descriptor) (err error) {
	return a.cli(d).Run(os.Args)
}
