package typapp

import (
	"github.com/urfave/cli/v2"
)

// Provider responsible to provide constructor [mock]
type Provider interface {
	Constructors() []*Constructor
}

// Preparer responsible to prepare the dependency[mock]
type Preparer interface {
	Preparations() []*Preparation
}

// Destroyer responsible to destroy the dependency [mock]
type Destroyer interface {
	Destructions() []*Destruction
}

// Commander responsible to return commands for App
type Commander interface {
	Commands(*Context) []*cli.Command
}
