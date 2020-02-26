package typrls

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Releaser responsible to release
type Releaser interface {
	Release(ctx context.Context, rls *Context) (err error)
}

// Context of release
type Context struct {
	*typcore.TypicalContext
	Alpha     bool
	Force     bool
	NoPublish bool
}
