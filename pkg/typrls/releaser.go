package typrls

import (
	"context"

	"github.com/typical-go/typical-go/pkg/git"
	"github.com/typical-go/typical-go/pkg/typbuild"
)

// Context of release
type Context struct {
	*typbuild.Context
	Name    string
	Tag     string
	GitLogs []*git.Log
	Alpha   bool
}

// Releaser responsible to release
type Releaser interface {
	Build(ctx context.Context, rls *Context) (binaries []string, err error)
	Publish(ctx context.Context, rls *Context, binaries []string) error
	Tag(ctx context.Context, version string, alpha bool) string
}
