package typbuild

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func (b *Build) cmdRun(c *Context) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the binary",
		SkipFlagParsing: true,
		Action: func(cliCtx *cli.Context) (err error) {
			ctx := cliCtx.Context
			if err = b.buildProject(ctx, c); err != nil {
				return
			}
			log.Info("Run the application")
			cmd := exec.CommandContext(ctx, typenv.AppBin, cliCtx.Args().Slice()...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			return cmd.Run()
		},
	}
}
