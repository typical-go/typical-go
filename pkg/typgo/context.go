package typgo

import (
	"context"
	"errors"
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
		mocker     *MockCommandRunner
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
			Stdout:  d.Stdout,
			Headers: LogHeaders(taskNames...),
		},
	}
}

// ExecuteCommand execute a command
func (c *Context) ExecuteCommand(basher Commander) error {
	bash := basher.Command()
	c.Logger.Command(bash)
	ctx := c.Ctx()
	if c.mocker != nil {
		return c.mocker.Run(bash)
	}
	return bash.ExecCmd(ctx).Run()
}

// ExecuteCommandLine execute bash command
func (c *Context) ExecuteCommandLine(commandLine string) error {
	if commandLine == "" {
		return errors.New("command line can't be empty")
	}
	return c.ExecuteCommand(CommandLine(commandLine))
}

// Ctx return golang context
func (c *Context) Ctx() context.Context {
	if c.Context == nil {
		return context.Background()
	}
	return c.Context.Context
}

// PatchBash typgo.RunBash for testing purpose
func (c *Context) PatchBash(mocks []*MockCommand) func(t *testing.T) {
	if c.mocker == nil {
		c.mocker = &MockCommandRunner{Mocks: mocks}
	}
	return func(t *testing.T) {
		require.NoError(t, c.mocker.Close())
	}
}
