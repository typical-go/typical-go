package typbuild

import (
	"context"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
)

func (b *Build) buildProject(ctx context.Context, c *Context) (err error) {
	if err = b.Prebuild(ctx, c); err != nil {
		return
	}
	log.Info("Build the application")
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
