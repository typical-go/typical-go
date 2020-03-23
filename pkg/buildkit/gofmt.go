package buildkit

import (
	"context"
	"os"
	"os/exec"
)

// GoFmt execute go fmt
type GoFmt struct {
	targets []string
	dir     string
}

// NewGoFmt return new instance of GoFmt
func NewGoFmt(targets ...string) *GoFmt {
	return &GoFmt{
		targets: targets,
	}
}

// WithDir return GoFmt with new dir
func (g *GoFmt) WithDir(dir string) *GoFmt {
	g.dir = dir
	return g
}

// Execute go fmt
func (g *GoFmt) Execute(ctx context.Context) error {
	args := []string{"fmt"}
	args = append(args, g.targets...)
	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Dir = g.dir
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
