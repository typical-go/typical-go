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
	Clean(*BuildContext) error
}

// Tester responsible to test the project
type Tester interface {
	Test(*BuildContext) error
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
	Precondition(c *BuildContext) error
}

// Runner responsible to run the project in local environment
type Runner interface {
	Run(c *BuildContext) error
}

// Context of buildtool
type Context struct {
	*typcore.Context
	*BuildTool
}

// BuildContext to create BuildContext
func (c *Context) BuildContext(cliCtx *cli.Context) *BuildContext {
	return &BuildContext{
		Context: c,
		Cli:     cliCtx,
	}
}

// BuildContext is context of build
type BuildContext struct {
	*Context
	Cli *cli.Context
}

// ReleaseContext is context of release
type ReleaseContext struct {
	*BuildContext
	Alpha   bool
	Tag     string
	GitLogs []*git.Log
}

// PublishContext is context of publish
type PublishContext struct {
	*ReleaseContext
	ReleaseFiles []string
}
