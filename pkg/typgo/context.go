package typgo

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/urfave/cli/v2"
)

type (
	// Context related with build task
	Context struct {
		*cli.Context
		Descriptor *Descriptor
		Stdout     io.Writer
	}
)

// NewContext return instance of context
func NewContext(c *cli.Context, d *Descriptor) *Context {
	return &Context{
		Context:    c,
		Descriptor: d,
		Stdout:     os.Stdout,
	}
}

// DummyContext return dummy context
func DummyContext() (*Context, *strings.Builder) {
	var out strings.Builder
	c := cli.NewContext(nil, &flag.FlagSet{}, nil)
	c.Command = &cli.Command{Name: "dummy"}
	context := &Context{
		Context: c,
		Descriptor: &Descriptor{
			ProjectName:    "some-project",
			ProjectVersion: "0.0.1",
		},
		Stdout: &out,
	}
	return context, &out
}

// Execute command
func (c *Context) Execute(basher Basher) error {
	bash := basher.Bash()
	c.printHeader()
	color.New(color.FgMagenta).Fprint(c.Stdout, "$ ")
	fmt.Fprintln(c.Stdout, bash)
	return RunBash(c.Ctx(), bash)
}

// ExecuteBash execute bash command
func (c *Context) ExecuteBash(commandLine string) error {
	if commandLine == "" {
		return errors.New("command line can't be empty")
	}
	slices := strings.Split(commandLine, " ")
	return c.Execute(&Bash{
		Name:   slices[0],
		Args:   slices[1:],
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
}

// Ctx return golang context
func (c *Context) Ctx() context.Context {
	return c.Context.Context
}

// Info log text
func (c *Context) Info(text string) {
	c.printHeader()
	fmt.Fprintln(c.Stdout, text)
}

// Infof formatted text
func (c *Context) Infof(format string, a ...interface{}) {
	c.printHeader()
	fmt.Fprintf(c.Stdout, format, a...)
}

func (c *Context) printHeader() {
	color.New(color.FgYellow).Fprint(c.Stdout, c.Descriptor.ProjectName)
	fmt.Fprint(c.Stdout, ":")
	color.New(color.FgBlue).Fprint(c.Stdout, c.Command.Name)
	fmt.Fprint(c.Stdout, "> ")
}
