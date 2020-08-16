package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

type (
	// CompileCmd compile command
	CompileCmd struct {
		Action
	}
	// StdCompile is standard compile
	StdCompile struct {
		MainPackage string // By default is "cmd/PROJECT_NAME"
		Output      string // By default is "bin/PROJECT_NAME"
		// By default is set variable typapp.Name to PROJECT_NAME
		// and typapp.Version to PROJECT-VERSION
		Ldflags fmt.Stringer
	}
)

//
// CompileCommand
//

var _ Cmd = (*CompileCmd)(nil)

// Command compile
func (c *CompileCmd) Command(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:    "compile",
		Aliases: []string{"c"},
		Usage:   "Compile the project",
		Action:  b.Action(c.Action),
	}
}

//
// StdCompile
//

var _ Action = (*StdCompile)(nil)

// Execute standard compile
func (s *StdCompile) Execute(c *Context) error {
	return c.Execute(&execkit.GoBuild{
		Output:      s.getOutput(c),
		MainPackage: s.getMainPackage(c),
		Ldflags:     s.getLdflags(c),
	})
}

func (s *StdCompile) getMainPackage(c *Context) string {
	if s.MainPackage == "" {
		s.MainPackage = fmt.Sprintf("./cmd/%s", c.BuildSys.ProjectName)
	}
	return s.MainPackage
}

func (s *StdCompile) getOutput(c *Context) string {
	if s.Output == "" {
		s.Output = fmt.Sprintf("bin/%s", c.BuildSys.ProjectName)
	}
	return s.Output
}

func (s *StdCompile) getLdflags(c *Context) fmt.Stringer {
	if s.Ldflags == nil {
		s.Ldflags = execkit.BuildVars{
			"github.com/typical-go/typical-go/pkg/typapp.Name":    c.BuildSys.ProjectName,
			"github.com/typical-go/typical-go/pkg/typapp.Version": c.BuildSys.ProjectVersion,
		}
	}
	return s.Ldflags
}
