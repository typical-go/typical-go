package typgo

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

type (
	// RunProject standard run
	RunProject struct {
		Before Action
		Binary string
	}
)

//
// RunCmd
//

var _ Tasker = (*RunProject)(nil)
var _ Action = (*RunProject)(nil)

// Task to run binary
func (r *RunProject) Task(sys *BuildSys) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project",
		SkipFlagParsing: true,
		Before:          sys.Action(r.Before),
		Action:          sys.Action(r),
	}
}

// Execute standard run
func (r *RunProject) Execute(c *Context) error {
	return c.Execute(&execkit.Command{
		Name:   r.getBinary(c),
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
}

func (r *RunProject) getBinary(c *Context) string {
	if r.Binary == "" {
		r.Binary = fmt.Sprintf("bin/%s", c.BuildSys.ProjectName)
	}
	return r.Binary
}
