package typcore

import (
	"github.com/urfave/cli/v2"
)

// Provider responsible to provide dependency
type Provider interface {
	Provide() []interface{}
}

// Preparer responsible to prepare
type Preparer interface {
	Prepare() []interface{}
}

// Destroyer responsible to destruct dependency
type Destroyer interface {
	Destroy() []interface{}
}

// Actionable responsible to provide action. Only available for AppModule
type Actionable interface {
	Action() interface{}
}

// Configurer responsible to create config
// `Prefix` is used by ConfigLoader to retrieve configuration value
// `Spec` (Specification) is used readme/env file generator. The value of spec will act as local environment value defined in .env file.
// `LoadFn` (Load Function) is required to provide in dependecies-injection container
type Configurer interface {
	Configure() (prefix string, spec interface{}, loadFn interface{})
}

// BuildCommander responsible to return commands for Build-Tool
type BuildCommander interface {
	BuildCommands(c *Context) []*cli.Command
}

// AppCommander responsible to return commands for App
type AppCommander interface {
	AppCommands(c *Context) []*cli.Command
}

// IsProvider return true if object implementation of provider
func IsProvider(obj interface{}) (ok bool) {
	_, ok = obj.(Provider)
	return
}

// IsPreparer return true obj implement Preparer
func IsPreparer(obj interface{}) (ok bool) {
	_, ok = obj.(Preparer)
	return
}

// IsDestroyer return true if object implementation of destructor
func IsDestroyer(obj interface{}) (ok bool) {
	_, ok = obj.(Destroyer)
	return
}

// IsActionable return true if object is actionable
func IsActionable(obj interface{}) bool {
	_, ok := obj.(Actionable)
	return ok
}

// IsConfigurer return true if object implementation of configurer
func IsConfigurer(obj interface{}) (ok bool) {
	_, ok = obj.(Configurer)
	return
}

// IsBuildCommander return true if obj implement commander
func IsBuildCommander(obj interface{}) (ok bool) {
	_, ok = obj.(BuildCommander)
	return
}

// IsAppCommander return true if object implementation of AppCLI
func IsAppCommander(obj interface{}) (ok bool) {
	_, ok = obj.(AppCommander)
	return
}
