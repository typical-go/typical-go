package typcore

import (
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

// NewApp return new instance of app
func NewApp(v interface{}) *App {
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
		a.WithPrepare(dep.(Preparer))
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
func (a *App) AppCommands(ac *AppContext) (cmds []*cli.Command) {
	for _, commander := range a.commanders {
		cmds = append(cmds, commander.AppCommands(ac)...)
	}
	return
}
