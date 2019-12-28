package typbuildtool

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func (t buildtool) cmdRun() *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the binary",
		SkipFlagParsing: true,
		Action:          t.runBinary,
	}
}

func (t buildtool) runBinary(c *cli.Context) (err error) {
	var (
		ctx = c.Context
	)
	if err = t.buildBinary(c); err != nil {
		return
	}
	log.Info("Run the application")
	args := []string(c.Args().Slice())
	cmd := exec.CommandContext(ctx, typenv.AppBin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
