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

	// DefaultConfigFile is default config file path
	DefaultConfigFile = "app.env"

	// DefaultPrecondition is default precondition flag
	DefaultPrecondition = true
)

var (
	_ typcore.App                 = (*App)(nil)
	_ typbuildtool.Preconditioner = (*App)(nil)
	_ typcfg.Configurer           = (*App)(nil)

	_ Provider  = (*App)(nil)
	_ Destroyer = (*App)(nil)
	_ Preparer  = (*App)(nil)
	_ Commander = (*App)(nil)
)

// App is typical application model
type App struct {
	appSources []string
	main       *typdep.Invocation
	modules    []interface{}
	configurer []typcfg.Configurer

	initFile     string
	configFile   string
	precondition bool
}

// EntryPoint create new instance of App with main invocation function
func EntryPoint(mainFn interface{}, appSource string, sources ...string) *App {
	return &App{
		appSources:   append([]string{appSource}, sources...),
		main:         typdep.NewInvocation(mainFn),
		initFile:     DefaultInitFile,
		precondition: DefaultPrecondition,
		configFile:   DefaultConfigFile,
	}
}

// Configures the application
func (a *App) Configures(configurer ...typcfg.Configurer) *App {
	a.configurer = configurer
	return a
}

// WithModules return app with appended module. Module should be implementation of Provider, Preparer (optional) and Destroyer (optional).
func (a *App) WithModules(modules ...interface{}) *App {
	a.modules = modules
	return a
}

// WithInitFile return app with new initFile
func (a *App) WithInitFile(initFile string) *App {
	a.initFile = initFile
	return a
}

// WithConfigFile return app with new configFile
func (a *App) WithConfigFile(configFile string) *App {
	a.configFile = configFile
	return a
}

// WithPrecondition return app with new precondition
func (a *App) WithPrecondition(precondition bool) *App {
	a.precondition = precondition
	return a
}

// Provide to return constructors
func (a *App) Provide() (constructors []*Constructor) {
	constructors = append(constructors, global...)
	for _, module := range a.modules {
		if provider, ok := module.(Provider); ok {
			constructors = append(constructors, provider.Provide()...)
		}
	}
	return
}

//Destroy to return destructor
func (a *App) Destroy() (destructors []*Destruction) {
	for _, module := range a.modules {
		if destroyer, ok := module.(Destroyer); ok {
			destructors = append(destructors, destroyer.Destroy()...)
		}
	}
	return
}

// Prepare to return preparations
func (a *App) Prepare() (preparations []*Preparation) {

	for _, module := range a.modules {
		if preparer, ok := module.(Preparer); ok {
			preparations = append(preparations, preparer.Prepare()...)
		}
	}
	return
}

// Commands to return commands
func (a *App) Commands(c *Context) (cmds []*cli.Command) {
	for _, module := range a.modules {
		if commander, ok := module.(Commander); ok {
			cmds = append(cmds, commander.Commands(c)...)
		}
	}
	return
}

// Configurations of app
func (a *App) Configurations() (cfgs []*typcfg.Configuration) {
	for _, module := range a.modules {
		if c, ok := module.(typcfg.Configurer); ok {
			cfgs = append(cfgs, c.Configurations()...)
		}
	}

	for _, c := range a.configurer {
		cfgs = append(cfgs, c.Configurations()...)
	}

	return
}

// AppSources return source for app
func (a *App) AppSources() []string {
	return a.appSources
}
