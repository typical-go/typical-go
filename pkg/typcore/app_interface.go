package typcore

import "github.com/urfave/cli/v2"

// AppInterface is interface of app
type AppInterface interface {
	EntryPointer
	Provider
	Preparer
	Destroyer
	AppCommander
}

// EntryPointer responsible to handle entry point
type EntryPointer interface{ EntryPoint() interface{} }

// Provider responsible to provide dependency
type Provider interface{ Provide() []interface{} }

// Preparer responsible to prepare
type Preparer interface{ Prepare() []interface{} }

// Destroyer responsible to destruct dependency
type Destroyer interface{ Destroy() []interface{} }

// AppCommander responsible to return commands for App
type AppCommander interface {
	AppCommands(*AppContext) []*cli.Command
}
