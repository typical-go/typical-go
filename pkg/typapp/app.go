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
)

// App is typical application model
type App struct {
	EntryPoint interface{}
	typcfg.Configurer
}

// Run the application
func (a *App) Run(d *typcore.Descriptor) (err error) {
	return createAppCli(a, d).Run(os.Args)
}

// Configurations of app
func (a *App) Configurations() []*typcfg.Configuration {
	if a.Configurer != nil {
		return a.Configurer.Configurations()
	}
	return nil
}

// Precondition the app
func (a *App) Precondition(c *typbuild.PrecondContext) (err error) {
	appPrecond := a.appPrecond(c)
	if appPrecond.NotEmpty() {
		c.AppendTemplate(appPrecond)
	}
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
