package typbuildtool

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/common/stdrun"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typbuild/prebld"
	"github.com/typical-go/typical-go/pkg/typbuildtool/internal/tmpl"

	log "github.com/sirupsen/logrus"
)

type standardPrebuilder struct{}

func (a *standardPrebuilder) Prebuild(ctx context.Context, c *typbuild.Context) (err error) {
	var constructors common.Strings
	if err = c.EachAnnotation("constructor", prebld.FunctionType, func(decl *prebld.Declaration, ann *prebld.Annotation) (err error) {
		constructors.Append(fmt.Sprintf("%s.%s", decl.File.Name, decl.SourceName))
		return
	}); err != nil {
		return
	}
	log.Info("Generate constructors")
	target := fmt.Sprintf("%s/%s/constructor_do_not_edit.go", c.Cmd, c.Name)
	if err = a.generateConstructor(ctx, target, c, constructors); err != nil {
		return
	}
	return
}

func (a *standardPrebuilder) generateConstructor(ctx context.Context, target string, c *typbuild.Context, constructors common.Strings) (err error) {
	defer common.ElapsedTimeFn("Generate constructor")()
	imports := []string{"github.com/typical-go/typical-go/pkg/typapp"}
	for _, dir := range c.Dirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", c.ModulePackage, dir))
		}
	}
	if err = stdrun.NewWriteTemplate(target, tmpl.Constructor, tmpl.ConstructorData{
		Imports:      imports,
		Constructors: constructors,
	}).Run(); err != nil {
		return
	}
	cmd := exec.CommandContext(ctx,
		fmt.Sprintf("%s/bin/goimports", build.Default.GOPATH),
		"-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
