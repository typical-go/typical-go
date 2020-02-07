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

// BuildInterface is interface of build
type BuildInterface interface {
	BuildCommander
	Prebuilder
	Validate() (err error)
	Releaser() Releaser
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

// Prebuilder responsible to prebuild task
type Prebuilder interface {
	Prebuild(ctx context.Context, bc *BuildContext) error
}

// AppCommander responsible to return commands for App
type AppCommander interface {
	AppCommands(*AppContext) []*cli.Command
}

// BuildCommander responsible to return commands for Build-Tool
type BuildCommander interface {
	BuildCommands(c *BuildContext) []*cli.Command
}
