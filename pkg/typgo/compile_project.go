package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

type (
	// CompileProject compile command
	CompileProject struct {
		MainPackage string // By default is "cmd/PROJECT_NAME"
		Output      string // By default is "bin/PROJECT_NAME"
		// By default is set variable typgo.AppName to PROJECT_NAME
		// and typgo.AppVersion to PROJECT-VERSION
		Ldflags fmt.Stringer
	}
)

var _ Cmd = (*CompileProject)(nil)
var _ Action = (*CompileProject)(nil)

// Command compile
func (p *CompileProject) Command(b *BuildSys) *cli.Command {
	return &cli.Command{
		Name:    "compile",
		Aliases: []string{"c"},
		Usage:   "Compile the project",
		Action:  b.Action(p),
	}
}

// Execute standard compile
func (p *CompileProject) Execute(c *Context) error {
	return c.Execute(&execkit.GoBuild{
		Output:      p.getOutput(c),
		MainPackage: p.getMainPackage(c),
		Ldflags:     p.getLdflags(c),
	})
}

func (p *CompileProject) getMainPackage(c *Context) string {
	if p.MainPackage == "" {
		p.MainPackage = fmt.Sprintf("./cmd/%s", c.BuildSys.AppName)
	}
	return p.MainPackage
}

func (p *CompileProject) getOutput(c *Context) string {
	if p.Output == "" {
		p.Output = fmt.Sprintf("bin/%s", c.BuildSys.AppName)
	}
	return p.Output
}

func (p *CompileProject) getLdflags(c *Context) fmt.Stringer {
	if p.Ldflags == nil {
		p.Ldflags = execkit.BuildVars{
			"github.com/typical-go/typical-go/pkg/typgo.AppName":    c.BuildSys.AppName,
			"github.com/typical-go/typical-go/pkg/typgo.AppVersion": c.BuildSys.AppVersion,
		}
	}
	return p.Ldflags
}
