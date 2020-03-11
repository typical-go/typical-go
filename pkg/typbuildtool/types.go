package typbuildtool

import (
	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/urfave/cli/v2"
)

// Builder reponsible to build
type Builder interface {
	Build(c *BuildContext) (dist BuildDistribution, err error)
}

// Cleaner responsible to clean the project
type Cleaner interface {
	Clean(*BuildContext) error
}

// Preconditioner responsible to precondition
type Preconditioner interface {
	Precondition(c *BuildContext) error
}

// Tester responsible to test the project
type Tester interface {
	Test(*BuildContext) error
}

// BuildDistribution is build distribution
type BuildDistribution interface {
	Run(*BuildContext) error
}

// Mocker responsible to mock
type Mocker interface {
	Mock(*MockContext) error
}

// Releaser responsible to release
type Releaser interface {
	Release(*ReleaseContext) (err error)
}

// Publisher reponsible to publish the release to external source
type Publisher interface {
	Publish(*PublishContext) error
}

// BuildContext is context of build
type BuildContext struct {
	*Context
	Cli *cli.Context
	Ast *typast.Ast
}

// MockContext is context of mock
type MockContext struct {
	*BuildContext
}

// ReleaseContext is context of release
type ReleaseContext struct {
	*BuildContext
	Alpha bool
}

// PublishContext is context of publish
type PublishContext struct {
	*ReleaseContext
	Tag      string
	Binaries []string
	GitLogs  []*git.Log
}
