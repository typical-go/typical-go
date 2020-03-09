package typbuildtool

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/runnerkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typbuildtool/internal/tmpl"

	log "github.com/sirupsen/logrus"
)

type stdPrebuilder struct{}

func (a *stdPrebuilder) Prebuild(c *BuildContext) (err error) {
	var constructors []string
	if err = c.Ast.EachAnnotation("constructor", typast.FunctionType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		constructors = append(constructors, fmt.Sprintf("%s.%s", decl.File.Name, decl.SourceName))
		return
	}); err != nil {
		return
	}
	log.Info("Generate constructors")
	target := fmt.Sprintf("%s/%s/constructor_do_not_edit.go", c.CmdFolder, c.Name)
	if err = a.generateConstructor(c, target, constructors); err != nil {
		return
	}
	return
}

func (a *stdPrebuilder) generateConstructor(c *BuildContext, target string, constructors []string) (err error) {
	defer common.ElapsedTimeFn("Generate constructor")()
	ctx := c.Cli.Context
	imports := []string{}
	for _, dir := range c.Dirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", c.ModulePackage, dir))
		}
	}
	if err = runnerkit.WriteTemplate(target, tmpl.Constructor, tmpl.ConstructorData{
		Imports:      imports,
		Constructors: constructors,
	}, 0666).Run(ctx); err != nil {
		return
	}
	cmd := exec.CommandContext(ctx,
		fmt.Sprintf("%s/bin/goimports", build.Default.GOPATH),
		"-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
