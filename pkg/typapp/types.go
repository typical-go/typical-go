package typapp

import "github.com/typical-go/typical-go/pkg/typdep"

// Dependency of app
type Dependency interface {
	Provider
	Destroyer
}

// EntryPointer responsible to handle entry point
type EntryPointer interface {
	EntryPoint() *typdep.Invocation
}

// Provider responsible to provide constructor
type Provider interface {
	Provide() []*typdep.Constructor
}

// Preparer responsible to prepare
type Preparer interface {
	Prepare() []*typdep.Invocation
}

// Destroyer responsible to destroy dependency
type Destroyer interface {
	Destroy() []*typdep.Invocation
}
