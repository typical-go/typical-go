package exor

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

// GoBuild builder
type GoBuild struct {
	ldflags []string
	out     string
	src     string
	stdout  io.Writer
	stderr  io.Writer
	stdin   io.Reader
}

// NewGoBuild return new instance of gobuild
func NewGoBuild(out, src string) *GoBuild {
	return &GoBuild{
		out: out,
		src: src,
	}
}

// SetVariable to set variable using linker
func (g *GoBuild) SetVariable(name string, value interface{}) *GoBuild {
	g.ldflags = append(g.ldflags, fmt.Sprintf("-X %s=%v", name, value))
	return g
}

// WithStdout return Command with new stdout
func (g *GoBuild) WithStdout(stdout io.Writer) *GoBuild {
	g.stdout = stdout
	return g
}

// WithStderr return Command with new stderr
func (g *GoBuild) WithStderr(stderr io.Writer) *GoBuild {
	g.stderr = stderr
	return g
}

// WithStdin return Command with new stdin
func (g *GoBuild) WithStdin(stdin io.Reader) *GoBuild {
	g.stdin = stdin
	return g
}

// Execute comand
func (g *GoBuild) Execute(ctx context.Context) (err error) {
	cmd := exec.CommandContext(ctx, "go", g.Args()...)
	cmd.Stdout = g.stdout
	cmd.Stderr = g.stderr
	cmd.Stdin = g.stdin
	return cmd.Run()
}

// Args is arguments for go build
func (g *GoBuild) Args() []string {
	args := []string{"build"}
	if len(g.ldflags) > 0 {
		args = append(args, "-ldflags", strings.Join(g.ldflags, " "))
	}
	args = append(args, "-o", g.out, g.src)
	return args
}
