package stdrun

import (
	"context"
	"os"
	"os/exec"
)

// GoFmt responsible to run go formatter
type GoFmt struct {
	targets []string
	dir     string
	ctx     context.Context
}

// NewGoFmt return new instaence of GoFmt
func NewGoFmt(ctx context.Context, targets ...string) *GoFmt {
	return &GoFmt{
		ctx:     ctx,
		targets: targets,
	}
}

// WithDir to set directory anr return Gofmt
func (g *GoFmt) WithDir(dir string) *GoFmt {
	g.dir = dir
	return g
}

// Run to making the directory
func (g *GoFmt) Run() error {
	args := []string{"fmt"}
	args = append(args, g.targets...)
	cmd := exec.CommandContext(g.ctx, "go", args...)
	cmd.Dir = g.dir
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
