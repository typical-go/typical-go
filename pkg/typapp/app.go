package typapp

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	_ typcore.App                 = (*App)(nil)
	_ typbuildtool.Preconditioner = (*App)(nil)
	_ typcfg.Configurer           = (*App)(nil)
	_ Provider                    = (*App)(nil)
	_ Destroyer                   = (*App)(nil)
	_ Preparer                    = (*App)(nil)
)

// App is typical application model
type App struct {
	appSources []string // TODO: remove this
	mainFn     interface{}
	imports    []interface{}
}

// EntryPoint create new instance of App with main invocation function
func EntryPoint(mainFn interface{}, appSource string, sources ...string) *App {
	return &App{
		appSources: append([]string{appSource}, sources...),
		mainFn:     mainFn,
	}
}

// Imports either Provider, Preparer, Destroyer or Configurations
func (a *App) Imports(imports ...interface{}) *App {
	a.imports = imports
	return a
}

// RunApp to run the applciation
func (a *App) RunApp(d *typcore.Descriptor) (err error) {
	return createAppCli(a, d).Run(os.Args)
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
	c.AppendTemplate(a.appPrecond(c))
	return
}

func (a *App) appPrecond(c *typbuildtool.PreconditionContext) *typtmpl.AppPrecond {
	appPrecond := typtmpl.NewAppPrecond()

	ctors, errs := typannot.GetConstructor(c.ASTStore())

	for _, ctor := range ctors {
		appPrecond.AppendCtor(ctor.Name, ctor.Def)
	}

	for _, err := range errs {
		c.Warnf("App-Precond: %s", err.Error())
	}

	for _, cfg := range a.Configurations() {
		appPrecond.AppendCfgCtor(cfg)
	}

	return appPrecond
}
