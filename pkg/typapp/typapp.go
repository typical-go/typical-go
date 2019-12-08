package typapp

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/typical-go/typical-go/pkg/typmodule"
	"github.com/typical-go/typical-go/pkg/utility/envfile"

	"os"

	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/urfave/cli/v2"
)

// Run the application
func Run(ctx *typctx.Context) {
	if err := ctx.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	appCli := &typcli.AppCli{
		Context: ctx,
	}
	app := cli.NewApp()
	app.Name = ctx.Name
	app.Usage = ""
	app.Description = ctx.Description
	app.Version = ctx.Version
	if actionable, ok := ctx.AppModule.(typmodule.Actionable); ok {
		app.Action = appCli.Action(actionable.Action())
	}
	app.Before = func(ctx *cli.Context) error {
		return envfile.Load()
	}
	if commander, ok := ctx.AppModule.(typcli.AppCommander); ok {
		app.Commands = commander.AppCommands(appCli)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
