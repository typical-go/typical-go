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
		Name    string   // By default is "run"
		Aliases []string // By default is "r"
		Usage   string   // By default is "Run the project"
		Before  Action
		Action
	}
	// StdRun standard run
	StdRun struct {
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
		Name:            r.getName(),
		Aliases:         r.getAliases(),
		Usage:           r.getUsage(),
		SkipFlagParsing: true,
		Before:          sys.ActionFn(r.Before),
		Action:          sys.ActionFn(r.Action),
	}
}

func (r *RunCmd) getName() string {
	if r.Name == "" {
		r.Name = "run"
	}
	return r.Name
}

func (r *RunCmd) getAliases() []string {
	if len(r.Aliases) < 1 {
		r.Aliases = []string{"r"}
	}
	return r.Aliases
}

func (r *RunCmd) getUsage() string {
	if r.Usage == "" {
		r.Usage = "Run the project"
	}
	return r.Usage
}

//
// StdRun
//

var _ Action = (*StdRun)(nil)

// Execute standard run
func (s *StdRun) Execute(c *Context) error {
	return c.Execute(&execkit.Command{
		Name:   s.getBinary(c),
		Args:   c.Args().Slice(),
		Stdout: os.Stdout,
		Stderr: os.Stderr,
		Stdin:  os.Stdin,
	})
}

func (s *StdRun) getBinary(c *Context) string {
	if s.Binary == "" {
		s.Binary = fmt.Sprintf("bin/%s", c.BuildSys.ProjectName)
	}
	return s.Binary
}
