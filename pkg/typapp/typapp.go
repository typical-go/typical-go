package typapp

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typobj"
	"github.com/typical-go/typical-go/pkg/utility/envfile"
	"github.com/urfave/cli/v2"
)

// Run the application
func Run(ctx *typctx.Context) {
	if err := ctx.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	appCli := typctx.NewCli(ctx, ctx.AppModule)
	app := cli.NewApp()
	app.Name = ctx.Name
	app.Usage = ""
	app.Description = ctx.Description
	app.Version = ctx.Version
	if actionable, ok := ctx.AppModule.(typobj.Actionable); ok {
		app.Action = appCli.PreparedAction(actionable.Action())
	}
	app.Before = func(ctx *cli.Context) error {
		return envfile.Load()
	}
	if commander, ok := ctx.AppModule.(typobj.AppCommander); ok {
		app.Commands = commander.AppCommands(appCli)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
