package typbuildtool

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Release this project
func (b *StdBuild) Release(c *ReleaseContext) (files []string, err error) {
	for _, target := range b.releaseTargets {
		c.Infof("Build release for %s", target)
		var binary string
		if binary, err = b.build(c.CliContext, c.Tag, target); err != nil {
			err = fmt.Errorf("Failed build release: %w", err)
			return
		}
		files = append(files, binary)
	}

	return
}

func (b *StdBuild) build(c *CliContext, tag string, target ReleaseTarget) (binary string, err error) {
	ctx := c.Cli.Context
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{c.Name, tag, goos, goarch}, "_")
	// TODO: Support CGO
	cmd := exec.CommandContext(ctx, "go", "build",
		"-o", fmt.Sprintf("%s/%s", b.releaseFolder, binary),
		"-ldflags", "-w -s",
		fmt.Sprintf("./%s/%s", c.BuildTool.cmdFolder, c.Name),
	)
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)
	cmd.Stdout = b.stdout
	cmd.Stderr = b.stderr
	if err = cmd.Run(); err != nil {
		return
	}
	return
}
