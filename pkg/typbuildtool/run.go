package typbuildtool

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/bash"
	"github.com/urfave/cli"
)

func (t buildtool) cmdRun() cli.Command {
	return cli.Command{
		Name:      "run",
		ShortName: "r",
		Usage:     "Run the binary",
		Flags: []cli.Flag{
			cli.BoolFlag{
				Name:  "no-build",
				Usage: "Run the binary without build",
			},
		},
		Action: t.runBinary,
	}
}

func (t buildtool) runBinary(ctx *cli.Context) (err error) {
	if !ctx.Bool("no-build") {
		if err = t.buildBinary(ctx); err != nil {
			return
		}
	}
	log.Info("Run the application")
	return bash.Run(typenv.App.BinPath, []string(ctx.Args())...)
}
