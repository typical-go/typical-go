package typapp

import (
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

// TypicalApp is typical application model
type TypicalApp struct {
	entryPoint     EntryPointer
	providers      []Provider
	preparers      []Preparer
	destroyers     []Destroyer
	commanders     []AppCommander
	projectSources []string
}

// AppCommander responsible to return commands for App
type AppCommander interface {
	AppCommands(*Context) []*cli.Command
}

// New return new instance of app
func New(v interface{}) *TypicalApp {
	app := &TypicalApp{
		projectSources: []string{common.PackageName(v)},
	}
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

// WithEntryPointer return app with new entry pointer
func (a *TypicalApp) WithEntryPointer(entryPoint EntryPointer) *TypicalApp {
	a.entryPoint = entryPoint
	return a
}

// WithSources return app with new source
func (a *TypicalApp) WithSources(sources ...string) *TypicalApp {
	return a
}

// AppendProvider return app with appended provider
func (a *TypicalApp) AppendProvider(provides ...Provider) *TypicalApp {
	a.providers = append(a.providers, provides...)
	return a
}

// AppendPreparer return app with appended preparer
func (a *TypicalApp) AppendPreparer(prepares ...Preparer) *TypicalApp {
	a.preparers = append(a.preparers, prepares...)
	return a
}

// AppendDestroyer return app with appended destroyer
func (a *TypicalApp) AppendDestroyer(destroys ...Destroyer) *TypicalApp {
	a.destroyers = append(a.destroyers, destroys...)
	return a
}

// AppendCommander return app with appended commander
func (a *TypicalApp) AppendCommander(commands ...AppCommander) *TypicalApp {
	a.commanders = append(a.commanders, commands...)
	return a
}

// AppendDependency return app with appended dependency
func (a *TypicalApp) AppendDependency(dependencies ...Dependency) *TypicalApp {
	for _, dep := range dependencies {
		a.AppendProvider(dep.(Provider))
		a.AppendDestroyer(dep.(Destroyer))
	}
	return a
}

// AppendProjectSource return app with appended project sources
func (a *TypicalApp) AppendProjectSource(sources ...string) *TypicalApp {
	a.projectSources = append(a.projectSources, sources...)
	return a
}

// EntryPoint of app
func (a *TypicalApp) EntryPoint() *typdep.Invocation {
	if a.entryPoint != nil {
		return a.entryPoint.EntryPoint()
	}
	return nil
}

// Provide to return constructors
func (a *TypicalApp) Provide() (constructors []*typdep.Constructor) {
	constructors = append(constructors, appConstructors...)
	for _, provider := range a.providers {
		constructors = append(constructors, provider.Provide()...)
	}
	return
}

//Destroy to return destructor
func (a *TypicalApp) Destroy() (destructors []*typdep.Invocation) {
	for _, destroyer := range a.destroyers {
		destructors = append(destructors, destroyer.Destroy()...)
	}
	return
}

// Prepare to return preparations
func (a *TypicalApp) Prepare() (preparations []*typdep.Invocation) {
	for _, preparer := range a.preparers {
		preparations = append(preparations, preparer.Prepare()...)
	}
	return
}

// AppCommands to return commands
func (a *TypicalApp) AppCommands(c *Context) (cmds []*cli.Command) {
	for _, commander := range a.commanders {
		cmds = append(cmds, commander.AppCommands(c)...)
	}
	return
}

// ProjectSources return source for app
func (a *TypicalApp) ProjectSources() []string {
	return a.projectSources
}

// Run application
func (a *TypicalApp) Run(d *typcore.Descriptor) (err error) {
	c := &Context{
		Descriptor: d,
		TypicalApp: a,
	}
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version
	app.Before = func(c *cli.Context) (err error) {
		if err = typcfg.LoadEnvFile(); err != nil {
			return
		}
		return
	}
	if entryPoint := a.EntryPoint(); entryPoint != nil {
		app.Action = c.ActionFunc(entryPoint)
	}
	app.Commands = a.AppCommands(c)
	return app.Run(os.Args)
}
