package application

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typicli"

	"os"

	"github.com/typical-go/typical-go/pkg/typictx"
	"github.com/urfave/cli"
)

// Run the application
func Run(ctx *typictx.Context) {
	ctxCli := &typicli.ContextCli{
		Context: ctx,
	}
	app := cli.NewApp()
	app.Name = ctx.Name
	app.Usage = ""
	app.Description = ctx.Description
	app.Version = ctx.Version
	app.Action = ctxCli.Action(ctx.AppModule.Run())
	app.Before = typicli.LoadEnvFile
	if commander, ok := ctx.AppModule.(typicli.AppCommander); ok {
		app.Commands = commander.AppCommands(ctxCli)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
