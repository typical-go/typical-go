package typapp

import (
	"github.com/urfave/cli/v2"
)

// Provider responsible to provide constructor [mock]
type Provider interface {
	Provide() []*Constructor
}

// Preparer responsible to prepare the dependency[mock]
type Preparer interface {
	Prepare() []*Preparation
}

// Destroyer responsible to destroy the dependency [mock]
type Destroyer interface {
	Destroy() []*Destruction
}

// Commander responsible to return commands for App
type Commander interface {
	Commands(*Context) []*cli.Command
}
