package execkit

import (
	"context"
	"fmt"
	"os"
	"time"
)

type (
	// GoTest builder
	GoTest struct {
		Targets      []string
		CoverProfile string
		Race         bool
		Timeout      time.Duration
	}
)

var _ Runner = (*GoTest)(nil)
var _ fmt.Stringer = (*GoTest)(nil)

// Command of go test
func (g *GoTest) Command() *Command {
	return &Command{
		Name:   "go",
		Args:   g.Args(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	}
}

// Args is argument for gotest
func (g *GoTest) Args() []string {
	args := []string{"test"}
	if g.Timeout > 0 {
		args = append(args, fmt.Sprintf("-timeout=%s", g.Timeout.String()))
	}
	if g.CoverProfile != "" {
		args = append(args, fmt.Sprintf("-coverprofile=%s", g.CoverProfile))
	}
	if g.Race {
		args = append(args, "-race")
	}
	return append(args, g.Targets...)
}

// Run gotest
func (g *GoTest) Run(ctx context.Context) error {
	return g.Command().Run(ctx)
}

func (g GoTest) String() string {
	return g.Command().String()
}
