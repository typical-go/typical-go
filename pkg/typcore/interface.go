package typcore

import (
	"context"

	"github.com/urfave/cli/v2"
)

// App is interface of app
type App interface {
	EntryPointer
	Provider
	Preparer
	Destroyer
	AppCommander
}

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

// ConfigurationInterface is interface of configuration
type ConfigurationInterface interface {
	Provider
	Loader() ConfigLoader
	ConfigMap() (keys []string, configMap ConfigMap)
}

// EntryPointer responsible to handle entry point
type EntryPointer interface {
	EntryPoint() interface{}
}

// Provider responsible to provide dependency
type Provider interface {
	Provide() []interface{}
}

// Preparer responsible to prepare
type Preparer interface {
	Prepare() []interface{}
}

// Destroyer responsible to destroy dependency
type Destroyer interface {
	Destroy() []interface{}
}

// AppCommander responsible to return commands for App
type AppCommander interface {
	AppCommands(*AppContext) []*cli.Command
}

// BuildCommander responsible to return commands for Build-Tool
type BuildCommander interface {
	BuildCommands(c *BuildContext) []*cli.Command
}
