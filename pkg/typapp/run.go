package typapp

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Run application
func (a *App) Run(d *typcore.Descriptor) (err error) {
	c := &Context{
		Descriptor: d,
		App:        a,
	}
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version
	app.Before = func(c *cli.Context) (err error) {
		if err = typcfg.LoadEnvFile(); err != nil {
			return
		}
		return
	}
	if entryPoint := a.EntryPoint(); entryPoint != nil {
		app.Action = c.ActionFunc(entryPoint)
	}
	app.Commands = a.AppCommands(c)
	return app.Run(os.Args)
}
