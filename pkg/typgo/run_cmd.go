package typgo

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

type (
	// RunCmd run command
	RunCmd struct {
		Before Action
		Action
	}
	// RunProject standard run
	RunProject struct {
		Binary string
	}
)

//
// RunCmd
//

var _ Cmd = (*RunCmd)(nil)

// Command run
func (r *RunCmd) Command(sys *BuildSys) *cli.Command {
	return &cli.Command{
		Name:            "run",
		Aliases:         []string{"r"},
		Usage:           "Run the project",
		SkipFlagParsing: true,
		Before:          sys.Action(r.Before),
		Action:          sys.Action(r.Action),
	}
}

//
// StdRun
//

var _ Action = (*RunProject)(nil)

// Execute standard run
func (s *RunProject) Execute(c *Context) error {
	return c.Execute(&execkit.Command{
		Name:   s.getBinary(c),
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
}

func (s *RunProject) getBinary(c *Context) string {
	if s.Binary == "" {
		s.Binary = fmt.Sprintf("bin/%s", c.BuildSys.AppName)
	}
	return s.Binary
}
