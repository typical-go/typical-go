package typbuildtool

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli"
)

func (t buildtool) cmdRun() cli.Command {
	return cli.Command{
		Name:            "run",
		ShortName:       "r",
		Usage:           "Run the binary",
		SkipFlagParsing: true,
		Action:          t.runBinary,
	}
}

func (t buildtool) runBinary(ctx *cli.Context) (err error) {
	if err = t.buildBinary(ctx); err != nil {
		return
	}
	log.Info("Run the application")
	args := []string(ctx.Args())
	cmd := exec.Command(typenv.App.BinPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
