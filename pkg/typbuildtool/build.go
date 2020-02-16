package typbuildtool

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typbuild"
)

func (b *Build) buildProject(ctx context.Context, c *typbuild.Context) (err error) {
	if err = b.prebuild(ctx, c); err != nil {
		return
	}
	log.Info("Build the application")
	cmd := exec.CommandContext(ctx,
		"go",
		"build",
		"-o", fmt.Sprintf("%s/%s", c.Bin, c.Name),
		fmt.Sprintf("./%s/%s", c.Cmd, c.Name),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (b *Build) prebuild(ctx context.Context, c *typbuild.Context) (err error) {
	for _, prebuilder := range b.prebuilders {
		if err = prebuilder.Prebuild(ctx, c); err != nil {
			return
		}
	}
	return
}
