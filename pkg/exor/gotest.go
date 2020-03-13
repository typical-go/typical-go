package exor

import (
	"context"
	"fmt"
	"io"
	"os/exec"
)

// GoTest builder
type GoTest struct {
	CoverProfile string
	Race         bool
	Targets      []string
	stdout       io.Writer
	stderr       io.Writer
	stdin        io.Reader
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

// WithStdout return Command with new stdout
func (g *GoTest) WithStdout(stdout io.Writer) *GoTest {
	g.stdout = stdout
	return g
}

// WithStderr return Command with new stderr
func (g *GoTest) WithStderr(stderr io.Writer) *GoTest {
	g.stderr = stderr
	return g
}

// WithStdin return Command with new stdin
func (g *GoTest) WithStdin(stdin io.Reader) *GoTest {
	g.stdin = stdin
	return g
}

// Execute comand
func (g *GoTest) Execute(ctx context.Context) (err error) {
	cmd := exec.CommandContext(ctx, "go", g.Args()...)
	cmd.Stdout = g.stdout
	cmd.Stderr = g.stderr
	cmd.Stdin = g.stdin
	return cmd.Run()
}

// Args is argument for gotest
func (g *GoTest) Args() []string {
	args := []string{"test"}
	if g.CoverProfile != "" {
		args = append(args, fmt.Sprintf("-coverprofile=%s", g.CoverProfile))
	}
	if g.Race {
		args = append(args, "-race")
	}
	args = append(args, g.Targets...)
	return args
}
