package typapp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typfactory"
)

// Precondition the app
func (a *App) Precondition(c *typbuildtool.CliContext) (err error) {
	c.Info("Precondition the typical-app")
	if err = a.generateConstructor(c, "typical/"+a.initFile); err != nil {
		return
	}
	return
}

func (a *App) generateConstructor(c *typbuildtool.CliContext, filename string) (err error) {
	var (
		pkgs         []string
		constructors []string
	)

	if constructors, err = walkForConstructor(c); err != nil {
		return
	}

	for _, cfg := range a.Configurations() {
		constructors = append(constructors, ConfigContructor(cfg))
	}

	for _, dir := range c.Core.AppDirs {
		if !strings.Contains(dir, "internal") {
			pkgs = append(pkgs, fmt.Sprintf("%s/%s", c.Core.ProjectPkg, dir))
		}
	}

	if err = typfactory.WriteFile(filename, 0777, &typfactory.InitialApp{
		Imports:      pkgs,
		Constructors: constructors,
	}); err != nil {
		return
	}

	return buildkit.NewGoImports(c.Core.TypicalTmp, filename).Execute(c.Context)
}

func walkForConstructor(c *typbuildtool.CliContext) (constructors []string, err error) {
	var (
		store *typast.ASTStore
	)

	if store, err = typast.Walk(c.Core.AppFiles...); err != nil {
		return
	}
	err = store.EachAnnotation("constructor", typast.FunctionType,
		func(decl *typast.Decl, ann *typast.Annotation) (err error) {
			constructors = append(constructors, fmt.Sprintf("%s.%s", decl.File.Name, decl.SourceName))
			return
		})
	return
}

// ConfigContructor is definition for  configuration constructor
func ConfigContructor(bean *typcfg.Configuration) string {
	typ := reflect.TypeOf(bean.Spec).String()
	tmpl := `func() (cfg %s, err error){
		cfg = new(%s)
		if err = typcfg.Process("%s", cfg); err != nil {
			return nil, err
		}
		return  
	}`
	return fmt.Sprintf(tmpl, typ, typ[1:], bean.Name)
}
