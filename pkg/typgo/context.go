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
	// DummyContextSetting option for dummy option
	DummyContextSetting struct {
		Output  *strings.Builder
		FlagSet *flag.FlagSet
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

// Execute command
func (c *Context) Execute(basher Basher) error {
	bash := basher.Bash()
	if c.Stdout != nil {
		c.printHeader()
		color.New(BashColor).Fprintln(c.Stdout, bash)
	}
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
	if c.Context == nil {
		return context.Background()
	}
	return c.Context.Context
}

// Info log text
func (c *Context) Info(text string) {
	if c.Stdout == nil {
		return
	}
	c.printHeader()
	fmt.Fprintln(c.Stdout, text)
}

// Infof formatted text
func (c *Context) Infof(format string, a ...interface{}) {
	if c.Stdout == nil {
		return
	}
	c.printHeader()
	fmt.Fprintf(c.Stdout, format, a...)
}

func (c *Context) printHeader() {
	if c.Descriptor != nil {
		color.New(ProjectNameColor).Fprint(c.Stdout, c.Descriptor.ProjectName)
	}
	if c.Context != nil && c.Command != nil {
		for _, name := range strings.Split(c.Command.FullName(), " ") {
			fmt.Fprint(c.Stdout, ":")
			color.New(TaskNameColor).Fprint(c.Stdout, name)
		}
	}
	fmt.Fprint(c.Stdout, "> ")
}

//
// DummyContextSetting
//

func (c *DummyContextSetting) String() string {
	return c.Output.String()
}
