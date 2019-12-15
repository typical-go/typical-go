package typcore

import "github.com/urfave/cli/v2"

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

// Actionable responsible to provide action
type Actionable interface {
	Action() interface{}
}

// Validator responsible to validate the struct
type Validator interface {
	Validate() error
}

// Configurer responsible to create config
type Configurer interface {
	Configure() (prefix string, spec interface{}, loadFn interface{})
}

// BuildCommander responsible to command
type BuildCommander interface {
	BuildCommands(c Cli) []*cli.Command
}

// AppCommander return command
type AppCommander interface {
	AppCommands(c Cli) []*cli.Command
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

// IsValidator return true if object is actionable
func IsValidator(obj interface{}) bool {
	_, ok := obj.(Validator)
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
