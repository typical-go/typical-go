package typbuildtool

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func (t buildtool) cmdTest() *cli.Command {
	return &cli.Command{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "Run the testing",
		Action:  t.runTesting,
	}
}

func (t buildtool) runTesting(ctx *cli.Context) error {
	log.Info("Run testings")
	targets := []string{
		"./app/...",
		"./pkg/...",
	}
	args := []string{"test", "-coverprofile=cover.out", "-race"}
	args = append(args, targets...)
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
