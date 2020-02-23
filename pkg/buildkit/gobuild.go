package buildkit

import (
	"context"
	"os/exec"
)

// GoBuild builder
type GoBuild struct {
	ldflags
	out string
	src string
}

// NewGoBuild return new instance of gobuild
func NewGoBuild(out, src string) *GoBuild {
	return &GoBuild{
		out: out,
		src: src,
	}
}

// Command to return exec.Cmd to execute the go build
func (g *GoBuild) Command(ctx context.Context) *exec.Cmd {
	args := []string{"build"}
	if len(g.ldflags.lines) > 0 {
		args = append(args, "-ldflags", g.ldflags.String())
	}
	args = append(args, "-o", g.out, g.src)

	return exec.CommandContext(ctx, "go", args...)
}
