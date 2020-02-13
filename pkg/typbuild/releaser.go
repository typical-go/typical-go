package typbuild

import "context"

// ReleaseContext is release context
type ReleaseContext struct {
	*Context
	Name       string
	Tag        string
	ChangeLogs []string
	Alpha      bool
}

// Releaser responsible to release
type Releaser interface {
	Build(ctx context.Context, rls *ReleaseContext) (binaries []string, err error)
	Publish(ctx context.Context, rls *ReleaseContext, binaries []string) error
	Tag(ctx context.Context, version string, alpha bool) (string, error)
}
