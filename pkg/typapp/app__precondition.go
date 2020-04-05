package typapp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typfactory"
)

// Precondition the app
func (a *App) Precondition(c *typbuildtool.BuildContext) (err error) {
	if !a.precondition {
		c.Info("Skip Precondition for typical-app")
		return
	}

	var constructors []string
	if err = c.Ast().EachAnnotation("constructor", typast.FunctionType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		constructors = append(constructors, fmt.Sprintf("%s.%s", decl.File.Name, decl.SourceName))
		return
	}); err != nil {
		return
	}

	if c.ConfigManager != nil {
		for _, bean := range c.Configurations() {
			constructors = append(constructors, ConfigContructor(bean))
		}
	}

	c.Info("Precondition the typical-app")
	if err = a.generateConstructor(c, "typical/"+a.initAppFilename, constructors); err != nil {
		return
	}
	return
}

func (a *App) generateConstructor(c *typbuildtool.BuildContext, filename string, constructors []string) (err error) {
	ctx := c.Cli.Context
	imports := []string{}

	for _, dir := range c.AppDirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", c.ProjectPackage, dir))
		}
	}

	if err = typfactory.WriteFile(filename, 0777, &typfactory.InitialApp{
		Imports:      imports,
		Constructors: constructors,
	}); err != nil {
		return
	}

	return buildkit.NewGoImports(c.TypicalTmp, filename).Execute(ctx)
}

// ConfigContructor is definition for  configuration constructor
func ConfigContructor(bean *typcore.Configuration) string {
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
