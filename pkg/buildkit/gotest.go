package buildkit

import (
	"fmt"
	"io"
	"time"

	"github.com/typical-go/typical-go/pkg/execkit"
)

// GoTest builder
type GoTest struct {
	Targets      []string
	CoverProfile string
	Race         bool
	Timeout      time.Duration
	Stdout       io.Writer
	Stderr       io.Writer
	Stdin        io.Reader
}

// Command of go test
func (g *GoTest) Command() *execkit.Command {
	return &execkit.Command{
		Name:   "go",
		Args:   g.Args(),
		Stdout: g.Stdout,
		Stderr: g.Stderr,
		Stdin:  g.Stdin,
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
