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

// TypicalApp is typical application model
type TypicalApp struct {
	appSources      []string
	appModule       interface{}
	modules         []interface{}
	initAppFilename string
	precondition    bool
}

// AppModule create new instance of TypicalApp with AppModule
func AppModule(appModule interface{}, appSources ...string) *TypicalApp {
	if len(appSources) < 1 {
		appSources = []string{common.PackageName(appModule)}
	}
	return &TypicalApp{
		appSources:      appSources,
		appModule:       appModule,
		initAppFilename: DefaultInitAppFilename,
		precondition:    DefaultPrecondition,
	}
}

// EntryPoint create new instance of TypicalApp with main invocation function
func EntryPoint(fn interface{}, appSources ...string) *TypicalApp {
	return &TypicalApp{
		appSources:      appSources,
		appModule:       NewMainInvocation(fn),
		initAppFilename: DefaultInitAppFilename,
		precondition:    DefaultPrecondition,
	}
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

// WithPrecondition return app with new precondition
func (a *TypicalApp) WithPrecondition(precondition bool) *TypicalApp {
	a.precondition = precondition
	return a
}

// EntryPoint of app
func (a *TypicalApp) EntryPoint() *MainInvocation {
	if entryPointer, ok := a.appModule.(EntryPointer); ok {
		return entryPointer.EntryPoint()
	}
	return nil
}

// Provide to return constructors
func (a *TypicalApp) Provide() (constructors []*Constructor) {
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
func (a *TypicalApp) Destroy() (destructors []*Destruction) {
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
func (a *TypicalApp) Prepare() (preparations []*Preparation) {
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
	return a.appSources
}
