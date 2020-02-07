package stdbuild

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

// CmdRun is command to run
func CmdRun(bc *typcore.BuildContext) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the binary",
		SkipFlagParsing: true,
		Action: func(c *cli.Context) (err error) {
			ctx := c.Context
			if err = buildProject(ctx, bc); err != nil {
				return
			}
			log.Info("Run the application")
			cmd := exec.CommandContext(ctx, typenv.AppBin, c.Args().Slice()...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Stdin = os.Stdin
			return cmd.Run()
		},
	}
}
