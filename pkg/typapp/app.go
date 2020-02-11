package typapp

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// App is application
type App struct {
	entryPoint EntryPointer
	providers  []Provider
	preparers  []Preparer
	destroyers []Destroyer
	commanders []AppCommander
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

// AppCommander responsible to return commands for App
type AppCommander interface {
	AppCommands(*typcore.AppContext) []*cli.Command
}

// New return new instance of app
func New(v interface{}) *App {
	app := new(App)
	if entryPoint, ok := v.(EntryPointer); ok {
		app.entryPoint = entryPoint
	}
	if provider, ok := v.(Provider); ok {
		app.providers = []Provider{provider}
	}
	if preparer, ok := v.(Preparer); ok {
		app.preparers = []Preparer{preparer}
	}
	if destroyer, ok := v.(Destroyer); ok {
		app.destroyers = []Destroyer{destroyer}
	}
	if commander, ok := v.(AppCommander); ok {
		app.commanders = []AppCommander{commander}
	}
	return app
}

// WithEntryPoint to set entry point
func (a *App) WithEntryPoint(entryPoint EntryPointer) *App {
	a.entryPoint = entryPoint
	return a
}

// WithProvide to set provide
func (a *App) WithProvide(provides ...Provider) *App {
	a.providers = append(a.providers, provides...)
	return a
}

// WithPrepare to set prepare
func (a *App) WithPrepare(prepares ...Preparer) *App {
	a.preparers = append(a.preparers, prepares...)
	return a
}

// WithDestroy to set destroy
func (a *App) WithDestroy(destroys ...Destroyer) *App {
	a.destroyers = append(a.destroyers, destroys...)
	return a
}

// WithCommand to set commanders
func (a *App) WithCommand(commands ...AppCommander) *App {
	a.commanders = append(a.commanders, commands...)
	return a
}

// WithDependency to set dependency
func (a *App) WithDependency(dependencies ...Dependency) *App {
	for _, dep := range dependencies {
		a.WithProvide(dep.(Provider))
		a.WithDestroy(dep.(Destroyer))
	}
	return a
}

// EntryPoint of app
func (a *App) EntryPoint() interface{} {
	if a.entryPoint != nil {
		return a.entryPoint.EntryPoint()
	}
	return nil
}

// Provide to return constructors
func (a *App) Provide() (constructors []interface{}) {
	constructors = append(constructors, appCtors...)
	for _, provider := range a.providers {
		constructors = append(constructors, provider.Provide()...)
	}
	return
}

//Destroy to return destructor
func (a *App) Destroy() (destructors []interface{}) {
	for _, destroyer := range a.destroyers {
		destructors = append(destructors, destroyer.Destroy()...)
	}
	return
}

// Prepare to return preparations
func (a *App) Prepare() (preparations []interface{}) {
	for _, preparer := range a.preparers {
		preparations = append(preparations, preparer.Prepare()...)
	}
	return
}

// AppCommands to return commands
func (a *App) AppCommands(ac *typcore.AppContext) (cmds []*cli.Command) {
	for _, commander := range a.commanders {
		cmds = append(cmds, commander.AppCommands(ac)...)
	}
	return
}
