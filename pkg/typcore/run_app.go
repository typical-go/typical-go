package typcore

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/urfave/cli/v2"
)

// RunApp the application
func RunApp(d *Descriptor) {
	var (
		actx *AppContext
		err  error
	)
	if actx, err = d.AppContext(); err != nil {
		log.Fatal(err.Error())
	}
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version
	app.Before = func(c *cli.Context) (err error) {
		if err = common.LoadEnvFile(); err != nil {
			return
		}
		return
	}
	if entryPoint := d.App.EntryPoint(); entryPoint != nil {
		app.Action = actx.ActionFunc(entryPoint)
	}
	for _, cmd := range d.App.AppCommands(actx) {
		app.Commands = append(app.Commands, cmd)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
