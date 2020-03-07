package buildkit

import (
	"context"
	"fmt"
	"os/exec"
)

// GoTest builder
type GoTest struct {
	CoverProfile string
	Race         bool
	Targets      []string
}

// NewGoTest return new instance of GoTest
func NewGoTest(targets ...string) *GoTest {
	return &GoTest{
		Targets: targets,
	}
}

// WithRace return GoTest with new race
func (g *GoTest) WithRace(race bool) *GoTest {
	g.Race = race
	return g
}

// WithCoverProfile return GoTest with new cover profile
func (g *GoTest) WithCoverProfile(coverProfile string) *GoTest {
	g.CoverProfile = coverProfile
	return g
}

// Command to return exec.Cmd to execute the go build
func (g *GoTest) Command(ctx context.Context) *exec.Cmd {
	args := []string{"test"}
	if g.CoverProfile != "" {
		args = append(args, fmt.Sprintf("-coverprofile=%s", g.CoverProfile))
	}
	if g.Race {
		args = append(args, "-race")
	}
	args = append(args, g.Targets...)

	return exec.CommandContext(ctx, "go", args...)
}
