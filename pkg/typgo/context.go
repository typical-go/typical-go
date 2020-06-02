package typgo

import (
	"context"
	"fmt"
	"os"

	"github.com/fatih/color"
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

	exec interface {
		Run(c context.Context) error
	}
)

// Execute command
func (c *Context) Execute(exec exec) error {
	w := os.Stdout
	color.New(color.FgMagenta).Fprint(w, "\n$ ")
	fmt.Fprintln(w, exec)

	return exec.Run(c.Ctx())
}
