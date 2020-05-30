package typgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/urfave/cli/v2"
)

type (
	// Context of build tool
	Context struct {
		typlog.Logger
		*cli.Context
		*BuildCli
	}
)

// Execute command
func (c *Context) Execute(cmd *execkit.Command) error {
	if cmd.Stderr == nil {
		cmd.Stderr = os.Stderr
	}

	if cmd.Stdin == nil {
		cmd.Stdin = os.Stdin
	}

	if cmd.Stdout == nil {
		cmd.Stdout = os.Stdout
	}

	execkit.PrintCommand(cmd, os.Stdout)

	return cmd.Run(c.Ctx())
}
