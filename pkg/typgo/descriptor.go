package typgo

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
)

var (
	_ typcore.AppLauncher       = (*Descriptor)(nil)
	_ typcore.BuildToolLauncher = (*Descriptor)(nil)

	_ Preconditioner = (*Descriptor)(nil)
)

type (

	// Descriptor describe the project
	Descriptor struct {

		// Name of the project (OPTIONAL).
		// It should be a characters with/without underscore or dash.
		// By default, project name is same with project folder
		Name string

		// Description of the project (OPTIONAL).
		Description string

		// Version of the project (OPTIONAL).
		// By default it is 0.0.1
		Version string

		Build

		Utility Utility

		Layouts []string

		SkipPrecond bool

		EntryPoint interface{}

		Configurer Configurer
	}
)

// LaunchApp to launch the app
func (d *Descriptor) LaunchApp() (err error) {
	if err = d.Validate(); err != nil {
		return
	}
	return launchApp(d)
}

// LaunchBuildTool to launch the build tool
func (d *Descriptor) LaunchBuildTool() (err error) {
	if err := d.Validate(); err != nil {
		return err
	}

	return launchBuildTool(d)
}

// Validate context
func (d *Descriptor) Validate() (err error) {
	if d.Version == "" {
		d.Version = "0.0.1"
	}

	if !ValidateName(d.Name) {
		return errors.New("Descriptor: bad name")
	}

	if err = common.Validate(d.Build); err != nil {
		return fmt.Errorf("Descriptor: %w", err)
	}

	return
}

// ValidateName to validate valid descriptor name
func ValidateName(name string) bool {
	if name == "" {
		return false
	}

	r, _ := regexp.Compile(`^[a-zA-Z\_\-]+$`)
	if !r.MatchString(name) {
		return false
	}
	return true
}

// Precondition for this project
func (d *Descriptor) Precondition(c *PrecondContext) (err error) {

	if d.Configurer != nil {
		if err = WriteConfig(typvar.ConfigFile, d.Configurer); err != nil {
			return
		}
	}

	if appPrecond := d.appPrecond(c); appPrecond.NotEmpty() {
		c.AppendTemplate(appPrecond)
	}

	LoadConfig(typvar.ConfigFile)
	return
}

func (d *Descriptor) appPrecond(c *PrecondContext) *typtmpl.AppPrecond {
	var (
		ctors    []*typtmpl.Ctor
		cfgCtors []*typtmpl.CfgCtor
		dtors    []*typtmpl.Dtor
	)

	ctorAnnots, errs := typannot.GetCtors(c.ASTStore)
	for _, a := range ctorAnnots {
		ctors = append(ctors, &typtmpl.Ctor{
			Name: a.Name,
			Def:  fmt.Sprintf("%s.%s", a.Decl.Pkg, a.Decl.Name),
		})
	}

	dtorAnnots, errs := typannot.GetDtors(c.ASTStore)
	for _, a := range dtorAnnots {
		dtors = append(dtors, &typtmpl.Dtor{
			Def: fmt.Sprintf("%s.%s", a.Decl.Pkg, a.Decl.Name),
		})
	}

	for _, err := range errs {
		c.Warnf("App-Precond: %s", err.Error())
	}

	if d.Configurer != nil {
		for _, cfg := range d.Configurer.Configurations() {
			specType := reflect.TypeOf(cfg.Spec).String()
			cfgCtors = append(cfgCtors, &typtmpl.CfgCtor{
				Name:      cfg.CtorName,
				Prefix:    cfg.Name,
				SpecType:  specType,
				SpecType2: specType[1:],
			})
		}
	}

	return &typtmpl.AppPrecond{
		Ctors:    ctors,
		CfgCtors: cfgCtors,
		Dtors:    dtors,
	}
}
