package typbuild

import (
	"context"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/urfave/cli/v2"
)

// Prebuilder responsible to prebuild task
type Prebuilder interface {
	Prebuild(ctx context.Context, c *Context) error
}

// BuildCommander responsible to return commands for Build-Tool
type BuildCommander interface {
	BuildCommands(c *Context) []*cli.Command
}

// ReleaseContext is release context
type ReleaseContext struct {
	*Context
	Name    string
	Tag     string
	GitLogs []*git.Log
	Alpha   bool
}

// Releaser responsible to release
type Releaser interface {
	Build(ctx context.Context, rls *ReleaseContext) (binaries []string, err error)
	Publish(ctx context.Context, rls *ReleaseContext, binaries []string) error
	Tag(ctx context.Context, version string, alpha bool) string
}
