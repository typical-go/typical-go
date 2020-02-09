package typcore

import (
	"github.com/urfave/cli/v2"
)

// App is interface of app
type App interface {
	EntryPointer
	Provider
	Preparer
	Destroyer
	AppCommander
	Run(*AppContext) error
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
