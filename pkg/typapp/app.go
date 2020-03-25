package typapp

import (
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

// TypicalApp is typical application model
type TypicalApp struct {
	projectSources  []string
	appModule       interface{}
	modules         []interface{}
	initAppFilename string
}

// Create new instance of app
func Create(appModule interface{}) *TypicalApp {
	app := &TypicalApp{
		projectSources:  []string{common.PackageName(appModule)},
		appModule:       appModule,
		initAppFilename: "init_app_do_not_edit.go",
	}
	return app
}

// WithProjectSources return app with new source
func (a *TypicalApp) WithProjectSources(sources ...string) *TypicalApp {
	a.projectSources = sources
	return a
}

// WithModules return app with appended module. Module should be implementation of Provider, Preparer (optional) and Destroyer (optional).
func (a *TypicalApp) WithModules(modules ...interface{}) *TypicalApp {
	a.modules = modules
	return a
}

// WithInitAppFilename return app with new initAppFilename
func (a *TypicalApp) WithInitAppFilename(initAppFilename string) *TypicalApp {
	a.initAppFilename = initAppFilename
	return a
}

// EntryPoint of app
func (a *TypicalApp) EntryPoint() *typdep.Invocation {
	if entryPointer, ok := a.appModule.(EntryPointer); ok {
		return entryPointer.EntryPoint()
	}
	return nil
}

// Provide to return constructors
func (a *TypicalApp) Provide() (constructors []*typdep.Constructor) {
	constructors = append(constructors, appConstructors...)
	if provider, ok := a.appModule.(Provider); ok {
		constructors = append(constructors, provider.Provide()...)
	}
	for _, module := range a.modules {
		if provider, ok := module.(Provider); ok {
			constructors = append(constructors, provider.Provide()...)
		}
	}
	return
}

//Destroy to return destructor
func (a *TypicalApp) Destroy() (destructors []*typdep.Invocation) {
	if destroyer, ok := a.appModule.(Destroyer); ok {
		destructors = append(destructors, destroyer.Destroy()...)
	}
	for _, module := range a.modules {
		if destroyer, ok := module.(Destroyer); ok {
			destructors = append(destructors, destroyer.Destroy()...)
		}
	}
	return
}

// Prepare to return preparations
func (a *TypicalApp) Prepare() (preparations []*typdep.Invocation) {
	if preparer, ok := a.appModule.(Preparer); ok {
		preparations = append(preparations, preparer.Prepare()...)
	}
	for _, module := range a.modules {
		if preparer, ok := module.(Preparer); ok {
			preparations = append(preparations, preparer.Prepare()...)
		}
	}
	return
}

// Commands to return commands
func (a *TypicalApp) Commands(c *Context) (cmds []*cli.Command) {
	if commander, ok := a.appModule.(Commander); ok {
		return commander.Commands(c)
	}
	return
}

// AppSources return source for app
func (a *TypicalApp) AppSources() []string {
	return a.projectSources
}
