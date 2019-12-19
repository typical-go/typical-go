package typapp

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/utility/envfile"
	"github.com/urfave/cli/v2"
)

// Run the application
func Run(descr *typcore.ProjectDescriptor) {
	if err := descr.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	appCli := typcore.NewCli(descr, descr.AppModule)
	app := cli.NewApp()
	app.Name = descr.Name
	app.Usage = ""
	app.Description = descr.Description
	app.Version = descr.Version
	if actionable, ok := descr.AppModule.(typcore.Actionable); ok {
		app.Action = appCli.PreparedAction(actionable.Action())
	}
	app.Before = func(ctx *cli.Context) error {
		return envfile.Load()
	}
	if commander, ok := descr.AppModule.(typcore.AppCommander); ok {
		app.Commands = commander.AppCommands(appCli)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
