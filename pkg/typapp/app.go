package typapp

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	_ typcore.Runner          = (*App)(nil)
	_ typbuild.Preconditioner = (*App)(nil)
	_ typcfg.Configurer       = (*App)(nil)
	_ Provider                = (*App)(nil)
	_ Destroyer               = (*App)(nil)
	_ Preparer                = (*App)(nil)
)

// App is typical application model
type App struct {
	EntryPoint interface{}
	Imports    []interface{}
}

// Run the application
func (a *App) Run(d *typcore.Descriptor) (err error) {
	return createAppCli(a, d).Run(os.Args)
}

// Constructors of app
func (a *App) Constructors() (constructors []*Constructor) {
	constructors = append(constructors, _ctors...)
	for _, module := range a.Imports {
		if provider, ok := module.(Provider); ok {
			constructors = append(constructors, provider.Constructors()...)
		}
	}
	return
}

// Destructors of app
func (a *App) Destructors() (destructors []*Destructor) {
	for _, module := range a.Imports {
		if destroyer, ok := module.(Destroyer); ok {
			destructors = append(destructors, destroyer.Destructors()...)
		}
	}
	return
}

// Preparations of app
func (a *App) Preparations() (preparations []*Preparation) {
	for _, module := range a.Imports {
		if preparer, ok := module.(Preparer); ok {
			preparations = append(preparations, preparer.Preparations()...)
		}
	}
	return
}

// Configurations of app
func (a *App) Configurations() (cfgs []*typcfg.Configuration) {
	for _, module := range a.Imports {
		if c, ok := module.(typcfg.Configurer); ok {
			cfgs = append(cfgs, c.Configurations()...)
		}
	}
	return
}

// Precondition the app
func (a *App) Precondition(c *typbuild.PrecondContext) (err error) {
	c.AppendTemplate(a.appPrecond(c))
	return
}

func (a *App) appPrecond(c *typbuild.PrecondContext) *typtmpl.AppPrecond {
	appPrecond := typtmpl.NewAppPrecond()
	ctors, errs := typannot.GetCtors(c.ASTStore())

	for _, ctor := range ctors {
		appPrecond.AppendCtor(ctor.Name, fmt.Sprintf("%s.%s", ctor.Decl.Pkg, ctor.Decl.Name))
	}

	for _, err := range errs {
		c.Warnf("App-Precond: %s", err.Error())
	}

	for _, cfg := range a.Configurations() {
		appPrecond.AppendCfgCtor(cfg)
	}

	return appPrecond
}
