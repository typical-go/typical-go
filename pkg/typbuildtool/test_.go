package typbuildtool

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/utility/bash"
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
	return bash.GoTest(t.TestTargets)
}
