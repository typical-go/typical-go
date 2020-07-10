package typgo

import (
	"strings"

	"github.com/typical-go/typical-go/pkg/git"
)

type (
	// Releaser responsible to release
	Releaser interface {
		Release(*ReleaseContext) error
	}
	// Releasers for composite release
	Releasers []Releaser
	// ReleaseContext contain data for release
	ReleaseContext struct {
		*Context
		Alpha   bool
		Tag     string
		GitLogs []*git.Log
	}
	// ReleaseFn release function
	ReleaseFn    func(*ReleaseContext) error
	releaserImpl struct {
		fn ReleaseFn
	}
)

//
// releaserImpl
//

// NewReleaser return new instance of Releaser
func NewReleaser(fn ReleaseFn) Releaser {
	return &releaserImpl{fn: fn}
}

func (r *releaserImpl) Release(c *ReleaseContext) error {
	return r.fn(c)
}

//
// Releaser
//

var _ Releaser = (Releasers)(nil)

// Release the releasers
func (r Releasers) Release(c *ReleaseContext) (err error) {
	for _, releaser := range r {
		if err = releaser.Release(c); err != nil {
			return
		}
	}
	return
}

//
// Command
//

func releaseTag(c *Context, alpha bool) string {
	version := "0.0.1"
	if c.Descriptor.Version != "" {
		version = c.Descriptor.Version
	}

	var builder strings.Builder
	builder.WriteString("v")
	builder.WriteString(version)
	// if c.BuildTool.IncludeBranch {
	// 	builder.WriteString("_")
	// 	builder.WriteString(git.Branch(c.Context))
	// }
	// if c.BuildTool.IncludeCommitID {
	// 	builder.WriteString("_")
	// 	builder.WriteString(git.LatestCommit(c.Context))
	// }
	if alpha {
		builder.WriteString("_alpha")
	}
	return builder.String()
}
