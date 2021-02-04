package typgo

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/oskit"
)

type (
	// RunBinary standard run
	RunBinary struct {
		Before Action
		Binary string
	}
)

var _ Tasker = (*RunBinary)(nil)
var _ Action = (*RunBinary)(nil)

// Task to run binary
func (r *RunBinary) Task() *Task {
	return &Task{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project",
		SkipFlagParsing: true,
		Before:          r.Before,
		Action:          r,
	}
}

// Execute standard run
func (r *RunBinary) Execute(c *Context) error {
	taskSignature(oskit.Stdout, "Run")
	return c.Execute(&Bash{
		Name:   r.getBinary(c),
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
}

func (r *RunBinary) getBinary(c *Context) string {
	if r.Binary == "" {
		r.Binary = fmt.Sprintf("bin/%s", c.Descriptor.ProjectName)
	}
	return r.Binary
}
