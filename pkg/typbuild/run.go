package typbuild

import (
	"context"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
)

func (b *Build) run(ctx context.Context, c *Context, args []string) (err error) {
	if err = b.buildProject(ctx, c); err != nil {
		return
	}
	log.Info("Run the application")
	cmd := exec.CommandContext(ctx, typenv.AppBin, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}
