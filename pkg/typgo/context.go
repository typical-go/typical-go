package typgo

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/fatih/color"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

var (
	// CtxExecWriter write for context execute
	CtxExecWriter io.Writer = os.Stdout
)

type (
	// Context of build tool
	Context struct {
		*cli.Context
		*BuildCli
	}
)

// Execute command
func (c *Context) Execute(runner execkit.Runner) error {
	return execute(c.Ctx(), runner)
}

// Ctx return golang context
func (c *Context) Ctx() context.Context {
	return c.Context.Context
}

func execute(ctx context.Context, runner execkit.Runner) error {
	color.New(color.FgMagenta).Fprint(CtxExecWriter, "\n$ ")
	fmt.Fprintln(CtxExecWriter, runner)
	return runner.Run(ctx)
}
