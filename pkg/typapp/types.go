package typapp

import (
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

// EntryPointer responsible to handle entry point
type EntryPointer interface {
	EntryPoint() *MainInvocation
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

// Commander responsible to return commands for App
type Commander interface {
	Commands(*Context) []*cli.Command
}
