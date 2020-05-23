package typgo

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
	"go.uber.org/dig"
)

var (
	_ typcore.AppLauncher   = (*Descriptor)(nil)
	_ typcore.BuildLauncher = (*Descriptor)(nil)

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
	if configFile := os.Getenv("CONFIG"); configFile != "" {
		_, err = LoadConfig(configFile)
	}

	di := dig.New()
	if err = setDependencies(di, d); err != nil {
		return
	}

	errs := common.GracefulRun(start(di, d), stop(di))
	return errs.Unwrap()
}

// LaunchBuild to launch the build tool
func (d *Descriptor) LaunchBuild() (err error) {
	if err := d.Validate(); err != nil {
		return err
	}

	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "Build-Tool"
	app.Description = d.Description
	app.Version = d.Version

	buildCli := createBuildCli(d)

	app.Before = beforeBuild(buildCli)
	app.Commands = buildCli.Commands()

	return app.Run(os.Args)
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
func (d *Descriptor) Precondition(c *Context) (err error) {

	if d.Configurer != nil {
		if err = WriteConfig(typvar.ConfigFile, d.Configurer); err != nil {
			return
		}
	}

	LoadConfig(typvar.ConfigFile)

	d.appPrecond(c)

	return
}

func (d *Descriptor) appPrecond(c *Context) {

	ctorAnnots, errs := typannot.GetCtors(c.ASTStore)
	for _, a := range ctorAnnots {
		c.Precond.Ctors = append(c.Precond.Ctors, &typtmpl.Ctor{
			Name: a.Name,
			Def:  fmt.Sprintf("%s.%s", a.Decl.Pkg, a.Decl.Name),
		})
	}

	dtorAnnots, errs := typannot.GetDtors(c.ASTStore)
	for _, a := range dtorAnnots {
		c.Precond.Dtors = append(c.Precond.Dtors, &typtmpl.Dtor{
			Def: fmt.Sprintf("%s.%s", a.Decl.Pkg, a.Decl.Name),
		})
	}

	for _, err := range errs {
		c.Warnf("App-Precond: %s", err.Error())
	}

	if d.Configurer != nil {
		for _, cfg := range d.Configurer.Configurations() {
			specType := reflect.TypeOf(cfg.Spec).String()
			c.Precond.CfgCtors = append(c.Precond.CfgCtors, &typtmpl.CfgCtor{
				Name:      cfg.Ctor,
				Prefix:    cfg.Name,
				SpecType:  specType,
				SpecType2: specType[1:],
			})
		}
	}
}
