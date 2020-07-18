package typgo

import (
	"context"
	"io"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/urfave/cli/v2"
)

var (
	// Stdout standard output
	Stdout io.Writer = os.Stdout
)

type (
	// Context of build tool
	Context struct {
		*cli.Context
		*Descriptor
		ASTStore *typast.ASTStore
		Imports  []string
	}
)

// Execute command
func (c *Context) Execute(cmder execkit.Commander) error {
	cmd := cmder.Command()
	cmd.Print(Stdout)
	return execkit.Run(c.Ctx(), cmd)
}

// Ctx return golang context
func (c *Context) Ctx() context.Context {
	return c.Context.Context
}
