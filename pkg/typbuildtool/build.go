package typbuildtool

import (
	"context"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typenv"

	"github.com/urfave/cli/v2"
)

func cmdBuild(d *typcore.ProjectDescriptor) *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the binary",
		Action: func(c *cli.Context) (err error) {
			ctx := c.Context
			log.Info("Build the application")
			return buildProject(ctx, d)
		},
	}
}

func buildProject(ctx context.Context, d *typcore.ProjectDescriptor) (err error) {
	if err = prebuild(ctx, d); err != nil {
		return
	}
	cmd := exec.CommandContext(ctx,
		"go",
		"build",
		"-o", typenv.AppBin,
		"./"+typenv.AppMainPath,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
