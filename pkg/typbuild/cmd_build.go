package typbuild

import (
	"context"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typenv"

	"github.com/urfave/cli/v2"
)

func (b *Build) cmdBuild(c *Context) *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the binary",
		Action: func(cliCtx *cli.Context) (err error) {
			log.Info("Build the application")
			return b.buildProject(cliCtx.Context, c)
		},
	}
}

func (b *Build) buildProject(ctx context.Context, c *Context) (err error) {
	if err = b.Prebuild(ctx, c); err != nil {
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
