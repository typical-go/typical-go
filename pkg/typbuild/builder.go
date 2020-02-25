package typbuild

import (
	"context"

	"github.com/urfave/cli/v2"
)

// Builder reponsible to build
type Builder interface {
	Build(ctx context.Context, c *Context) error
}

// Prebuilder responsible to prebuild
type Prebuilder interface {
	Prebuild(ctx context.Context, c *Context) error
}

// BuildCommander responsible to return commands for Build-Tool
type BuildCommander interface {
	BuildCommands(c *Context) []*cli.Command
}
