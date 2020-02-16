package stdrls

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/typical-go/typical-go/pkg/typbuild"
)

// Build the distribution
func (r *Releaser) Build(ctx context.Context, rel *typbuild.ReleaseContext) (binaries []string, err error) {
	for _, target := range r.targets {
		var binary string
		if binary, err = r.build(ctx, rel, target); err != nil {
			return
		}
		binaries = append(binaries, binary)
	}
	return
}

func (r *Releaser) build(ctx context.Context, rel *typbuild.ReleaseContext, target Target) (binary string, err error) {
	goos := target.OS()
	goarch := target.Arch()
	binary = strings.Join([]string{rel.Name, rel.Tag, goos, goarch}, "_")
	// TODO: Support CGO
	cmd := exec.CommandContext(ctx, "go", "build",
		"-o", fmt.Sprintf("%s/%s", rel.Release, binary),
		"-ldflags", "-w -s",
		fmt.Sprintf("./%s/%s", rel.Cmd, rel.Name),
	)
	cmd.Env = append(os.Environ(), "GOOS="+goos, "GOARCH="+goarch)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stdout
	if err = cmd.Run(); err != nil {
		return
	}
	return
}
