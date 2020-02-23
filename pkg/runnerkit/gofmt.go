package runnerkit

import (
	"context"
	"os"
	"os/exec"
)

// GoFmt runner
func GoFmt(targets ...string) Runner {
	return &goFmt{
		targets: targets,
	}
}

// GoFmtWithDir runner
func GoFmtWithDir(dir string, targets ...string) Runner {
	return &goFmt{
		targets: targets,
		dir:     dir,
	}
}

type goFmt struct {
	targets []string
	dir     string
}

func (g *goFmt) Run(ctx context.Context) error {
	args := []string{"fmt"}
	args = append(args, g.targets...)
	cmd := exec.CommandContext(ctx, "go", args...)
	cmd.Dir = g.dir
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
