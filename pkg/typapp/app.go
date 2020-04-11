package typapp

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typdep"
	"github.com/urfave/cli/v2"
)

const (
	// DefaultInitFile is default init file path
	DefaultInitFile = "init_app_do_not_edit.go"
)

var (
	_ typcore.App                 = (*App)(nil)
	_ typbuildtool.Preconditioner = (*App)(nil)
	_ typcfg.Configurer           = (*App)(nil)
	_ Provider                    = (*App)(nil)
	_ Destroyer                   = (*App)(nil)
	_ Preparer                    = (*App)(nil)
	_ Commander                   = (*App)(nil)
)

// App is typical application model
type App struct {
	appSources []string
	main       *typdep.Invocation
	imports    []interface{}

	initFile string
}

// EntryPoint create new instance of App with main invocation function
func EntryPoint(mainFn interface{}, appSource string, sources ...string) *App {
	return &App{
		appSources: append([]string{appSource}, sources...),
		main:       typdep.NewInvocation(mainFn),
		initFile:   DefaultInitFile,
	}
}

// Imports either Provider, Preparer, Destroyer or Configurations
func (a *App) Imports(imports ...interface{}) *App {
	a.imports = imports
	return a
}

// InitFile define path to generate initial requirement
func (a *App) InitFile(initFile string) *App {
	a.initFile = initFile
	return a
}

// Constructors of app
func (a *App) Constructors() (constructors []*Constructor) {
	constructors = append(constructors, global...)
	for _, module := range a.imports {
		if provider, ok := module.(Provider); ok {
			constructors = append(constructors, provider.Constructors()...)
		}
	}
	return
}

// Destructions of app
func (a *App) Destructions() (destructors []*Destruction) {
	for _, module := range a.imports {
		if destroyer, ok := module.(Destroyer); ok {
			destructors = append(destructors, destroyer.Destructions()...)
		}
	}
	return
}

// Preparations of app
func (a *App) Preparations() (preparations []*Preparation) {
	for _, module := range a.imports {
		if preparer, ok := module.(Preparer); ok {
			preparations = append(preparations, preparer.Preparations()...)
		}
	}
	return
}

// Commands to return commands
func (a *App) Commands(c *Context) (cmds []*cli.Command) {
	for _, module := range a.imports {
		if commander, ok := module.(Commander); ok {
			cmds = append(cmds, commander.Commands(c)...)
		}
	}
	return
}

// Configurations of app
func (a *App) Configurations() (cfgs []*typcfg.Configuration) {
	for _, module := range a.imports {
		if c, ok := module.(typcfg.Configurer); ok {
			cfgs = append(cfgs, c.Configurations()...)
		}
	}
	return
}

// AppSources return source for app
func (a *App) AppSources() []string {
	return a.appSources
}
