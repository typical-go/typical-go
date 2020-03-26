package buildkit

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

// GoTest builder
type GoTest struct {
	targets      []string
	coverProfile string
	race         bool
	timeout      time.Duration

	stdout io.Writer
	stderr io.Writer
	stdin  io.Reader
}

// NewGoTest return new instance of GoTest
func NewGoTest(targets ...string) *GoTest {
	return &GoTest{
		targets: targets,
		timeout: 20 * time.Second,
		stdout:  os.Stdout,
		stderr:  os.Stderr,
	}
}

// WithTimeout return GotTest with new timeout
func (g *GoTest) WithTimeout(timeout time.Duration) *GoTest {
	g.timeout = timeout
	return g
}

// WithRace return GoTest with new race
func (g *GoTest) WithRace(race bool) *GoTest {
	g.race = race
	return g
}

// WithCoverProfile return GoTest with new cover profile
func (g *GoTest) WithCoverProfile(coverProfile string) *GoTest {
	g.coverProfile = coverProfile
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
	if len(g.targets) < 1 {
		return errors.New("Nothing to test")
	}
	cmd := exec.CommandContext(ctx, "go", g.Args()...)
	cmd.Stdout = g.stdout
	cmd.Stderr = g.stderr
	cmd.Stdin = g.stdin
	return cmd.Run()
}

// Args is argument for gotest
func (g *GoTest) Args() []string {
	args := []string{"test"}

	if g.timeout > 0 {
		args = append(args, fmt.Sprintf("-timeout=%s", g.timeout.String()))
	}

	if g.coverProfile != "" {
		args = append(args, fmt.Sprintf("-coverprofile=%s", g.coverProfile))
	}

	if g.race {
		args = append(args, "-race")
	}

	return append(args, g.targets...)
}
