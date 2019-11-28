package typbuildtool

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"

	"github.com/urfave/cli"
)

func (t buildtool) cmdBuild() cli.Command {
	return cli.Command{
		Name:      "build",
		ShortName: "b",
		Usage:     "Build the binary",
		Action:    t.buildBinary,
	}
}

func (t buildtool) buildBinary(ctx *cli.Context) error {
	log.Info("Build the application")
	cmd := exec.Command("go", "build",
		"-o", typenv.AppBin(t.Name),
		"./"+typenv.AppMainPkg(t.Name),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
