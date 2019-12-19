package typapp

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/utility/common"
	"github.com/urfave/cli/v2"
)

// Run the application
func Run(d *typcore.ProjectDescriptor) {
	if err := d.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	ctx := typcore.NewContext(d)
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = ""
	app.Description = d.Description
	app.Version = d.Version
	if actionable, ok := d.AppModule.(typcore.Actionable); ok {
		app.Action = ctx.PreparedAction(actionable.Action())
	}
	app.Before = func(ctx *cli.Context) error {
		return common.LoadEnvFile()
	}
	if commander, ok := d.AppModule.(typcore.AppCommander); ok {
		app.Commands = commander.AppCommands(ctx)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
