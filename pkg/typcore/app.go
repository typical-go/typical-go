package typcore

import (
	"github.com/urfave/cli/v2"
)

// App is application
type App struct {
	entryPoint EntryPointer
	provides   []Provider
	prepares   []Preparer
	destroys   []Destroyer
	commands   []AppCommander
}

// NewApp return new instance of app
func NewApp() *App {
	return &App{}
}

// WithEntryPoint to set entry point
func (a *App) WithEntryPoint(entryPoint EntryPointer) *App {
	a.entryPoint = entryPoint
	return a
}

// WithProvide to set provide
func (a *App) WithProvide(provides ...Provider) *App {
	a.provides = append(a.provides, provides...)
	return a
}

// WithPrepare to set prepare
func (a *App) WithPrepare(prepares ...Preparer) *App {
	a.prepares = append(a.prepares, prepares...)
	return a
}

// WithDestroy to set destroy
func (a *App) WithDestroy(destroys ...Destroyer) *App {
	a.destroys = append(a.destroys, destroys...)
	return a
}

// WithCommand to set commanders
func (a *App) WithCommand(commands ...AppCommander) *App {
	a.commands = append(a.commands, commands...)
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
	for _, provider := range a.provides {
		constructors = append(constructors, provider.Provide()...)
	}
	return
}

//Destroy to return destructor
func (a *App) Destroy() (destructors []interface{}) {
	for _, destroyer := range a.destroys {
		destructors = append(destructors, destroyer.Destroy()...)
	}
	return
}

// Prepare to return preparations
func (a *App) Prepare() (preparations []interface{}) {
	for _, preparer := range a.prepares {
		preparations = append(preparations, preparer.Prepare()...)
	}
	return
}

// AppCommands to return commands
func (a *App) AppCommands(ac *AppContext) (cmds []*cli.Command) {
	for _, commander := range a.commands {
		cmds = append(cmds, commander.AppCommands(ac)...)
	}
	return
}
