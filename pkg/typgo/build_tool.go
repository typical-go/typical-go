package typgo

import (
	"strings"

	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/urfave/cli/v2"
)

type (
	// BuildTool detail
	BuildTool struct {
		*Descriptor
	}

	// Context of build tool
	Context struct {
		typlog.Logger

		Cli *cli.Context
		*BuildTool
	}

	// CliFunc is command line function
	CliFunc func(*Context) error
)

// ActionFunc to return related action func
func (c *BuildTool) ActionFunc(name string, fn CliFunc) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		return fn(&Context{
			Logger: typlog.Logger{
				Name: strings.ToUpper(name),
			},
			Cli:       cli,
			BuildTool: c,
		})
	}
}

// Commands to return command
func (c *BuildTool) Commands() (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdTest(c),
		cmdRun(c),
		cmdPublish(c),
		cmdClean(c),
	}

	if c.Utility != nil {
		for _, cmd := range c.Utility.Commands(c) {
			cmds = append(cmds, cmd)
		}
	}

	return cmds
}
