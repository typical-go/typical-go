package typapp

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Run the application
func Run(d *typcore.ProjectDescriptor) {
	appCtx := typcore.NewAppContext(d)
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version
	app.Before = func(c *cli.Context) (err error) {
		if err = d.Validate(); err != nil {
			return
		}
		if err = common.LoadEnvFile(); err != nil {
			return
		}
		return
	}
	if entryPoint := d.App.EntryPoint(); entryPoint != nil {
		app.Action = appCtx.ActionFunc(entryPoint)
	}
	for _, cmd := range d.App.AppCommands(appCtx) {
		app.Commands = append(app.Commands, cmd)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
