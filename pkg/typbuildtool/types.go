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
	*BuildTool
	Core *typcore.Context
}

// CliFunc is command line function
type CliFunc func(*CliContext) error

// ActionFunc to return related action func
func (c *Context) ActionFunc(fn CliFunc) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		return fn(&CliContext{
			Context:   cli,
			Core:      c.Core,
			BuildTool: c.BuildTool,
		})
	}
}

// CliContext is context of build
type CliContext struct {
	*cli.Context
	Core      *typcore.Context
	BuildTool *BuildTool
}

// Info logger
func (c *CliContext) Info(args ...interface{}) {
	c.Core.Info(args...)
}

// Infof logger
func (c *CliContext) Infof(format string, args ...interface{}) {
	c.Core.Infof(format, args)
}

// Warn logger
func (c *CliContext) Warn(args ...interface{}) {
	c.Core.Warn(args...)
}

// Warnf logger
func (c *CliContext) Warnf(format string, args ...interface{}) {
	c.Core.Warnf(format, args...)
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
