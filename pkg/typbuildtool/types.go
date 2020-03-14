package typbuildtool

import (
	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// Builder reponsible to build
type Builder interface {
	Build(c *Context) (dist BuildDistribution, err error)
}

// Cleaner responsible to clean the project
type Cleaner interface {
	Clean(*Context) error
}

// Tester responsible to test the project
type Tester interface {
	Test(*Context) error
}

// BuildDistribution is build distribution
type BuildDistribution interface {
	Run(*Context) error
}

// Mocker responsible to mock
type Mocker interface {
	Mock(*Context) error
}

// Releaser responsible to release
type Releaser interface {
	Release(*ReleaseContext) (files []string, err error)
}

// Publisher reponsible to publish the release to external source
type Publisher interface {
	Publish(*PublishContext) error
}

// Context of buildtool
type Context struct {
	*typcore.Context
	Cli *cli.Context
}

// NewContext return new instance of Context
func NewContext(tc *typcore.Context, cc *cli.Context) *Context {
	return &Context{
		Context: tc,
		Cli:     cc,
	}
}

// ReleaseContext is context of release
type ReleaseContext struct {
	*Context
	Alpha   bool
	Tag     string
	GitLogs []*git.Log
}

// PublishContext is context of publish
type PublishContext struct {
	*ReleaseContext
	ReleaseFiles []string
}
