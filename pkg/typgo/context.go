package typgo

import (
	"context"
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"
)

type (
	// Context related with build task
	Context struct {
		*cli.Context
		Logger
		Descriptor *Descriptor
		mocker     *BashMocker
	}
)

// NewContext return instance of context
func NewContext(c *cli.Context, d *Descriptor) *Context {
	var taskNames []string
	if c.Command != nil {
		taskNames = strings.Fields(c.Command.FullName())
	}
	return &Context{
		Context:    c,
		Descriptor: d,
		Logger: Logger{
			Stdout:      d.Stdout,
			ProjectName: d.ProjectName,
			TaskNames:   taskNames,
		},
	}
}

// Execute command
func (c *Context) Execute(basher Basher) error {
	bash := basher.Bash()
	c.Logger.Bash(bash)
	ctx := c.Ctx()
	if c.mocker != nil {
		return c.mocker.Run(bash)
	}
	return bash.ExecCmd(ctx).Run()
}

// ExecuteBash execute bash command
func (c *Context) ExecuteBash(commandLine string) error {
	if commandLine == "" {
		return errors.New("command line can't be empty")
	}
	slices := strings.Fields(commandLine)
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

// PatchBash typgo.RunBash for testing purpose
func (c *Context) PatchBash(mocks []*MockBash) func(t *testing.T) {
	if c.mocker == nil {
		c.mocker = &BashMocker{Mocks: mocks}
	}
	return func(t *testing.T) {
		require.NoError(t, c.mocker.Close())
	}
}
