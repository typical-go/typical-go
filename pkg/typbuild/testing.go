package typbuild

import (
	"context"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

func (b *Build) test(ctx context.Context, c *Context) error {
	log.Info("Run testings")
	targets := []string{
		"./app/...",
		"./pkg/...",
	}
	args := []string{"test", "-coverprofile=cover.out", "-race"}
	args = append(args, targets...)
	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
