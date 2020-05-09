package typcore

import (
	"strings"

	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/urfave/cli/v2"
)

// Context of buildtool
type Context struct {
	*BuildTool
	*Descriptor

	AppDirs  []string
	AppFiles []string
}

// CliContext is context of build
type CliContext struct {
	typlog.Logger

	Cli       *cli.Context
	Name      string
	Core      *Context
	BuildTool *BuildTool
}

// CliFunc is command line function
type CliFunc func(*CliContext) error

// ActionFunc to return related action func
func (c *Context) ActionFunc(name string, fn CliFunc) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		return fn(&CliContext{
			Name: name,
			Logger: typlog.Logger{
				Name: strings.ToUpper(name),
			},
			Cli:       cli,
			Core:      c,
			BuildTool: c.BuildTool,
		})
	}
}
