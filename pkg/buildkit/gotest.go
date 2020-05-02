package buildkit

import (
	"fmt"
	"time"
)

// GoTest builder
type GoTest struct {
	Targets      []string
	CoverProfile string
	Race         bool
	Timeout      time.Duration
}

// Command of go test
func (g *GoTest) Command() *Command {
	return &Command{
		Name: "go",
		Args: g.Args(),
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
