package typapp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/typical-go/typical-go/pkg/typcore"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typfactory"
)

const (
	constructorTag = "constructor"
)

// Precondition the app
func (a *App) Precondition(c *typbuildtool.PreconditionContext) (err error) {
	c.Info("Precondition the typical-app")
	if err = a.generateConstructor(c, "typical/"+a.initFile); err != nil {
		return
	}
	return
}

func (a *App) generateConstructor(c *typbuildtool.PreconditionContext, filename string) (err error) {
	var (
		constructors []string
	)

	store := c.ASTStore()

	for _, a := range store.Annots {
		if a.Equal(constructorTag, typast.Function) {
			constructors = append(constructors, ctorDef(a))
		}
	}

	for _, cfg := range a.Configurations() {
		constructors = append(constructors, cfgCtorDef(cfg))
	}

	if err = typfactory.WriteFile(filename, 0777, &typfactory.InitialApp{
		Imports:      retrImports(c.Core),
		Constructors: constructors,
	}); err != nil {
		return
	}

	return buildkit.NewGoImports(c.Core.TypicalTmp, filename).Execute(c.Ctx)
}

func retrImports(c *typcore.Context) (imports []string) {
	for _, dir := range c.AppDirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, importDef(c, dir))
		}
	}
	return
}

func importDef(c *typcore.Context, dir string) string {
	return fmt.Sprintf("%s/%s", c.ProjectPkg, dir)
}

func ctorDef(a *typast.Annotation) string {
	return fmt.Sprintf("%s.%s", a.File.Name, a.SourceName)
}

func cfgCtorDef(bean *typcfg.Configuration) string {
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
