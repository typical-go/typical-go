package typgo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

func launchBuild(d *Descriptor) (err error) {
	if err := d.Validate(); err != nil {
		return err
	}

	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "Build-Tool"
	app.Description = d.Description
	app.Version = d.Version

	buildCli := createBuildCli(d)

	app.Before = buildCli.ActionFn("BEFORE_BUILD", beforeBuild)
	app.After = buildCli.ActionFn("AFTER_BUILD", afterBuild)
	app.Commands = commands(buildCli)

	return app.Run(os.Args)
}

func beforeBuild(c *Context) (err error) {
	if c.Configurer != nil {
		if err = WriteConfig(typvar.ConfigFile, c.Configurer); err != nil {
			return
		}
	}

	LoadConfig(typvar.ConfigFile)

	if !c.Descriptor.SkipPrecond {
		if err = precond(c); err != nil {
			return
		}
	}
	return
}

func afterBuild(c *Context) (err error) {
	store := c.BuildCli.ASTStore
	b, _ := json.MarshalIndent(store.Annots, "", "\t")
	if err = ioutil.WriteFile(fmt.Sprintf("%s/annots.json", typvar.TypicalTmp), b, 0777); err != nil {
		return
	}
	return
}

func commands(b *BuildCli) (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdTest(b),
		cmdRun(b),
		cmdPublish(b),
		cmdClean(b),
	}

	if b.Utility != nil {
		for _, cmd := range b.Utility.Commands(b) {
			cmds = append(cmds, cmd)
		}
	}

	return cmds
}

func precond(c *Context) (err error) {

	if err = appPrecond(c); err != nil {
		return
	}

	path := typvar.Precond(c.Descriptor.Name)
	os.Remove(path)

	if c.Precond.NotEmpty() {
		if err = typtmpl.WriteFile(path, 0777, c.Precond); err != nil {
			return
		}
		if err = buildkit.GoImports(c.Ctx(), path); err != nil {
			return
		}
	}
	return
}

func appPrecond(c *Context) (err error) {

	ctorAnnots, errs := typannot.GetCtors(c.ASTStore)
	for _, a := range ctorAnnots {
		c.Precond.Ctors = append(c.Precond.Ctors, &typtmpl.Ctor{
			Name: a.Param.Name,
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

	cfgr := c.Descriptor.Configurer

	if cfgr != nil {
		for _, cfg := range cfgr.Configurations() {
			specType := reflect.TypeOf(cfg.Spec).String()
			c.Precond.CfgCtors = append(c.Precond.CfgCtors, &typtmpl.CfgCtor{
				Name:      cfg.Ctor,
				Prefix:    cfg.Name,
				SpecType:  specType,
				SpecType2: specType[1:],
			})
		}
	}

	return
}
