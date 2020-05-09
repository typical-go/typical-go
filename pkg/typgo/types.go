package typgo

import (
	"github.com/typical-go/typical-go/pkg/git"
)

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

// Runner2 responsible to run the project in local environment
type Runner2 interface {
	Run(c *CliContext) error
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
