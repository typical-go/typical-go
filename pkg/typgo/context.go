package typgo

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/urfave/cli/v2"
)

var (
	// CtxExecWriter write for context execute
	CtxExecWriter io.Writer = os.Stdout
)

type (
	// Context of build tool
	Context struct {
		typlog.Logger
		*cli.Context
		*BuildCli
	}
	exec interface {
		Run(c context.Context) error
	}
)

// Execute command
func (c *Context) Execute(exec exec) error {
	color.New(color.FgMagenta).Fprint(CtxExecWriter, "\n$ ")
	fmt.Fprintln(CtxExecWriter, exec)
	return exec.Run(c.Ctx())
}

// Ctx return golang context
func (c *Context) Ctx() context.Context {
	return c.Context.Context
}
