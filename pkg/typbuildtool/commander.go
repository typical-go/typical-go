package typbuildtool

import (
	"github.com/urfave/cli/v2"
)

// BuildCommander responsible to return commands for Build-Tool
type BuildCommander interface {
	BuildCommands(c *Context) []*cli.Command
}
