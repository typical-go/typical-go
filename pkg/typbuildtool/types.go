package typbuildtool

import (
	"github.com/urfave/cli/v2"
)

// Commander responsible to return commands for Build-Tool
type Commander interface {
	Commands(c *Context) []*cli.Command
}
