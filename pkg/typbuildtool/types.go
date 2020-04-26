package typbuildtool

import (
	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Utility of build-tool
type Utility interface {
	Commands(c *Context) []*cli.Command
}

// Cleaner responsible to clean the project
type Cleaner interface {
	Clean(*CliContext) error
}

// Tester responsible to test the project
type Tester interface {
	Test(*CliContext) error
}

// Releaser responsible to release
type Releaser interface {
	Release(*ReleaseContext) (files []string, err error)
}

// Publisher responsible to publish the release to external source
type Publisher interface {
	Publish(*PublishContext) error
}

// Preconditioner responsible to precondition
type Preconditioner interface {
	Precondition(c *CliContext) error
}

// Runner responsible to run the project in local environment
type Runner interface {
	Run(c *CliContext) error
}

// Context of buildtool
type Context struct {
	*typcore.Context
	BuildTool *BuildTool
}

// CliContext to create CliContext
func (c *Context) CliContext(cli *cli.Context) *CliContext {
	return &CliContext{
		Context: c,
		Cli:     cli,
	}
}

// CliContext is context of build
type CliContext struct {
	*Context
	Cli *cli.Context
}

// ReleaseContext is context of release
type ReleaseContext struct {
	*CliContext
	Alpha   bool
	Tag     string
	GitLogs []*git.Log
}

// PublishContext is context of publish
type PublishContext struct {
	*ReleaseContext
	ReleaseFiles []string
}
