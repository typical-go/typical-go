package typbuildtool

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typbuild"
)

func (b *BuildTool) buildProject(ctx context.Context, c *typbuild.Context) (err error) {
	if err = b.prebuild(ctx, c); err != nil {
		return
	}
	log.Info("Build the application")
	cmd := exec.CommandContext(ctx,
		"go",
		"build",
		"-o", fmt.Sprintf("%s/%s", c.BinFolder, c.Name),
		fmt.Sprintf("./%s/%s", c.CmdFolder, c.Name),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (b *BuildTool) prebuild(ctx context.Context, c *typbuild.Context) (err error) {
	for _, prebuilder := range b.prebuilders {
		if err = prebuilder.Prebuild(ctx, c); err != nil {
			return
		}
	}
	return
}
