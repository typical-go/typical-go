package typapp

import (
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

// EntryPointer responsible to handle entry point
type EntryPointer interface {
	EntryPoint() *typdep.Invocation
}

// Provider responsible to provide constructor [mock]
type Provider interface {
	Provide() []*typdep.Constructor
}

// Preparer responsible to prepare [mock]
type Preparer interface {
	Prepare() []*typdep.Invocation
}

// Destroyer responsible to destroy dependency [mock]
type Destroyer interface {
	Destroy() []*typdep.Invocation
}

// AppCommander responsible to return commands for App
type AppCommander interface {
	AppCommands(*Context) []*cli.Command
}

// Module to be imported
type Module interface {
	Provider
	Preparer
	Destroyer
}
