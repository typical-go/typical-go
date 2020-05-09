package typapp

import (
	"fmt"
	"os"
	"reflect"

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
func (a *App) Constructors() []*Constructor {
	ctors := _ctors
	for _, module := range a.Imports {
		if provider, ok := module.(Provider); ok {
			ctors = append(ctors, provider.Constructors()...)
		}
	}
	return ctors
}

// Destructors of app
func (a *App) Destructors() []*Destructor {
	dtors := _dtors
	for _, module := range a.Imports {
		if destroyer, ok := module.(Destroyer); ok {
			dtors = append(dtors, destroyer.Destructors()...)
		}
	}
	return dtors
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
	var (
		ctors    []*typtmpl.Ctor
		cfgCtors []*typtmpl.CfgCtor
		dtors    []*typtmpl.Dtor
	)

	store := c.ASTStore()

	ctorAnnots, errs := typannot.GetCtors(store)
	for _, a := range ctorAnnots {
		ctors = append(ctors, &typtmpl.Ctor{
			Name: a.Name,
			Def:  fmt.Sprintf("%s.%s", a.Decl.Pkg, a.Decl.Name),
		})
	}

	dtorAnnots, errs := typannot.GetDtors(store)
	for _, a := range dtorAnnots {
		dtors = append(dtors, &typtmpl.Dtor{
			Def: fmt.Sprintf("%s.%s", a.Decl.Pkg, a.Decl.Name),
		})
	}

	for _, err := range errs {
		c.Warnf("App-Precond: %s", err.Error())
	}

	for _, cfg := range a.Configurations() {
		specType := reflect.TypeOf(cfg.Spec).String()
		cfgCtors = append(cfgCtors, &typtmpl.CfgCtor{
			Name:      cfg.CtorName,
			Prefix:    cfg.Name,
			SpecType:  specType,
			SpecType2: specType[1:],
		})
	}

	return &typtmpl.AppPrecond{
		Ctors:    ctors,
		CfgCtors: cfgCtors,
		Dtors:    dtors,
	}
}
