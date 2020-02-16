package typbuildtool

import (
	"context"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typbuild"
)

func (b *Build) test(ctx context.Context, c *typbuild.Context) error {
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
