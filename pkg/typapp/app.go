package typapp

import (
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/urfave/cli/v2"
)

const (
	// DefaultInitAppFilename is default value for init-app filename
	DefaultInitAppFilename = "init_app_do_not_edit.go"

	// DefaultPrecondition is default value for precondition flag
	DefaultPrecondition = true
)

// App is typical application model
type App struct {
	appSources []string
	appModule  interface{}
	modules    []interface{}

	initAppFilename string
	precondition    bool
}

// AppModule create new instance of App with AppModule
func AppModule(appModule interface{}, appSources ...string) *App {
	if len(appSources) < 1 {
		appSources = []string{common.PackageName(appModule)}
	}
	return &App{
		appSources:      appSources,
		appModule:       appModule,
		initAppFilename: DefaultInitAppFilename,
		precondition:    DefaultPrecondition,
	}
}

// EntryPoint create new instance of App with main invocation function
func EntryPoint(fn interface{}, appSource string, sources ...string) *App {
	return &App{
		appSources:      append([]string{appSource}, sources...),
		appModule:       NewMainInvocation(fn),
		initAppFilename: DefaultInitAppFilename,
		precondition:    DefaultPrecondition,
	}
}

// WithModules return app with appended module. Module should be implementation of Provider, Preparer (optional) and Destroyer (optional).
func (a *App) WithModules(modules ...interface{}) *App {
	a.modules = modules
	return a
}

// WithInitAppFilename return app with new initAppFilename
func (a *App) WithInitAppFilename(initAppFilename string) *App {
	a.initAppFilename = initAppFilename
	return a
}

// WithPrecondition return app with new precondition
func (a *App) WithPrecondition(precondition bool) *App {
	a.precondition = precondition
	return a
}

// EntryPoint of app
func (a *App) EntryPoint() *MainInvocation {
	if entryPointer, ok := a.appModule.(EntryPointer); ok {
		return entryPointer.EntryPoint()
	}
	return nil
}

// Provide to return constructors
func (a *App) Provide() (constructors []*Constructor) {
	constructors = append(constructors, global...)
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
func (a *App) Destroy() (destructors []*Destruction) {
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
func (a *App) Prepare() (preparations []*Preparation) {
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
func (a *App) Commands(c *Context) (cmds []*cli.Command) {
	if commander, ok := a.appModule.(Commander); ok {
		cmds = append(cmds, commander.Commands(c)...)
	}
	for _, module := range a.modules {
		if commander, ok := module.(Commander); ok {
			cmds = append(cmds, commander.Commands(c)...)
		}
	}
	return
}

// AppSources return source for app
func (a *App) AppSources() []string {
	return a.appSources
}
