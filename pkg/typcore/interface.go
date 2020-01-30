package typcore

import (
	"context"

	"github.com/urfave/cli/v2"
)

// AppInterface is interface of app
type AppInterface interface {
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

// Dependency of app
type Dependency interface {
	Provider
	Destroyer
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
