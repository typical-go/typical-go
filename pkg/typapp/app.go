package typapp

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/urfave/cli/v2"
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
	main       *Invocation
	imports    []interface{}
}

// EntryPoint create new instance of App with main invocation function
func EntryPoint(mainFn interface{}, appSource string, sources ...string) *App {
	return &App{
		appSources: append([]string{appSource}, sources...),
		main:       NewInvocation(mainFn),
	}
}

// RunApp to run the applciation
func (a *App) RunApp(d *typcore.Descriptor) (err error) {
	return createAppCli(a, d).Run(os.Args)
}

// Imports either Provider, Preparer, Destroyer or Configurations
func (a *App) Imports(imports ...interface{}) *App {
	a.imports = imports
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

// Destructors of app
func (a *App) Destructors() (destructors []*Destructor) {
	for _, module := range a.imports {
		if destroyer, ok := module.(Destroyer); ok {
			destructors = append(destructors, destroyer.Destructors()...)
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

// Precondition the app
func (a *App) Precondition(c *typbuildtool.PreconditionContext) (err error) {
	c.Info("Precondition the typical-app")

	provideCtor := typtmpl.NewProvideCtor()

	for _, a := range GetCtorAnnot(c) {
		provideCtor.AppendCtor(a.Name, a.Def)
	}

	for _, cfg := range a.Configurations() {
		provideCtor.AppendCfgCtor("", cfg)
	}

	// c.AppendImport(retrImports(c.Core)...)
	c.AppendWriter(provideCtor)

	return
}
