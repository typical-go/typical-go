package typapp

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/typcore"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
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

	store := c.ASTStore()

	provideCtor := &typfactory.ProvideCtor{}

	for _, a := range store.Annots {
		if a.Equal(constructorTag, typast.Function) {
			provideCtor.FnDefs = append(provideCtor.FnDefs, ctorDef(a))
		}
	}

	for _, cfg := range a.Configurations() {
		provideCtor.Cfgs = append(provideCtor.Cfgs, cfg)
	}

	initial := typfactory.NewInitialApp(retrImports(c.Core)...)
	initial.AppendWithWriter(provideCtor)

	if err = typfactory.WriteFile(filename, 0777, initial); err != nil {
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
	return fmt.Sprintf("%s.%s", a.Pkg, a.Name)
}
