package typbuildtool

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func (t buildtool) cmdTest() cli.Command {
	return cli.Command{
		Name:      "test",
		ShortName: "t",
		Usage:     "Run the testing",
		Action:    t.runTesting,
	}
}

func (t buildtool) runTesting(ctx *cli.Context) error {
	log.Info("Run testings")
	args := []string{"test"}
	args = append(args, t.TestTargets...)
	args = append(args,
		"-coverprofile=cover.out",
		"-race")
	cmd := exec.Command("go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
