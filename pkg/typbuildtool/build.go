package typbuildtool

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/utility/bash"

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
	return bash.GoBuild(typenv.App.BinPath, typenv.App.SrcPath)
}
