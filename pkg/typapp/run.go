package typapp

import (
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Run application
func (a *App) Run(actx *typcore.AppContext) (err error) {
	app := cli.NewApp()
	app.Name = actx.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = actx.Description
	app.Version = actx.Version
	app.Before = func(c *cli.Context) (err error) {
		if err = common.LoadEnvFile(); err != nil {
			return
		}
		return
	}
	if entryPoint := a.EntryPoint(); entryPoint != nil {
		app.Action = actx.ActionFunc(entryPoint)
	}
	app.Commands = a.AppCommands(actx)
	return app.Run(os.Args)
}
