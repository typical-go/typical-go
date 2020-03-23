package typapp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typapp/internal/tmpl"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Precondition the app
func (a *TypicalApp) Precondition(c *typbuildtool.BuildContext) (err error) {
	var constructors []string

	if err = c.Ast().EachAnnotation("constructor", typast.FunctionType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		constructors = append(constructors, fmt.Sprintf("%s.%s", decl.File.Name, decl.SourceName))
		return
	}); err != nil {
		return
	}

	if c.ConfigManager != nil {
		for _, bean := range c.Configurations() {
			constructors = append(constructors, configDefinition(bean))
		}
	}

	c.Info("Generate constructors")
	target := "typical/init_app_do_not_edit.go"
	if err = a.generateConstructor(c, target, constructors); err != nil {
		return
	}
	return
}

func configDefinition(bean *typcore.Configuration) string {
	typ := reflect.TypeOf(bean.Spec()).String()
	return fmt.Sprintf(`func(cfgMngr typcore.ConfigManager) (%s, error){
		cfg, err := cfgMngr.RetrieveConfig("%s")
		if err != nil {
			return nil, err
		}
		return  cfg.(%s), nil 
	}`, typ, bean.Name(), typ)
}

func (a *TypicalApp) generateConstructor(c *typbuildtool.BuildContext, filename string, constructors []string) (err error) {
	ctx := c.Cli.Context
	imports := []string{}

	for _, dir := range c.ProjectDirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", c.ProjectPackage, dir))
		}
	}

	if err = common.WriteTemplate(filename, tmpl.Constructor, tmpl.ConstructorData{
		Imports:      imports,
		Constructors: constructors,
	}); err != nil {
		return
	}

	return typcore.GoImport(ctx, c.Context.Context, filename)
}
