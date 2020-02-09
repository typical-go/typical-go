package typcore

import "context"

// Build is interface of build
type Build interface {
	BuildCommander
	Prebuilder
	Validate() (err error)
	Releaser() Releaser
}

// Prebuilder responsible to prebuild task
type Prebuilder interface {
	Prebuild(ctx context.Context, bc *BuildContext) error
}

// Releaser responsible to release
type Releaser interface {
	BuildRelease(ctx context.Context, name, tag string, changeLogs []string, alpha bool) (binaries []string, err error)
	Publish(ctx context.Context, name, tag string, changeLogs, binaries []string, alpha bool) (err error)
	Tag(ctx context.Context, version string, alpha bool) (tag string, err error)
	Validate() error
}
