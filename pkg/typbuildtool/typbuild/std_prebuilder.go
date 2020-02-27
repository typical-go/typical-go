package typbuild

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/runnerkit"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typbuildtool/typbuild/internal/tmpl"

	log "github.com/sirupsen/logrus"
)

type stdPrebuilder struct{}

func (a *stdPrebuilder) Prebuild(ctx context.Context, c *Context) (err error) {
	var constructors common.Strings
	if err = c.EachAnnotation("constructor", typast.FunctionType, func(decl *typast.Declaration, ann *typast.Annotation) (err error) {
		constructors.Append(fmt.Sprintf("%s.%s", decl.File.Name, decl.SourceName))
		return
	}); err != nil {
		return
	}
	log.Info("Generate constructors")
	target := fmt.Sprintf("%s/%s/constructor_do_not_edit.go", c.CmdFolder, c.Name)
	if err = a.generateConstructor(ctx, target, c, constructors); err != nil {
		return
	}
	return
}

func (a *stdPrebuilder) generateConstructor(ctx context.Context, target string, c *Context, constructors common.Strings) (err error) {
	defer common.ElapsedTimeFn("Generate constructor")()
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
