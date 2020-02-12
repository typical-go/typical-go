package typbuild

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/runn/stdrun"
	"github.com/typical-go/typical-go/pkg/typbuild/internal/tmpl"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcore/walker"
	"github.com/typical-go/typical-go/pkg/typenv"
)

type standardPrebuilder struct{}

func (a *standardPrebuilder) Prebuild(ctx context.Context, c *Context) (err error) {
	var constructors common.Strings
	if err = c.EachAnnotation("constructor", walker.FunctionType, func(decl *walker.Declaration, ann *walker.Annotation) (err error) {
		constructors.Append(fmt.Sprintf("%s.%s", decl.File.Name, decl.SourceName))
		return
	}); err != nil {
		return
	}
	log.Info("Generate constructors")
	if err = a.generateConstructor(ctx, typenv.GeneratedConstructor, c, constructors); err != nil {
		return
	}
	return
}

func (a *standardPrebuilder) generateConstructor(ctx context.Context, target string, c *Context, constructors common.Strings) (err error) {
	defer common.ElapsedTimeFn("Generate constructor")()
	imports := []string{"github.com/typical-go/typical-go/pkg/typapp"}
	for _, dir := range c.Dirs {
		imports = append(imports, fmt.Sprintf("%s/%s", c.Package, dir))
	}
	data := tmpl.ConstructorData{
		Imports:      imports,
		Constructors: constructors,
	}
	if err = stdrun.NewWriteTemplate(target, tmpl.Constructor, data).Run(); err != nil {
		return
	}
	cmd := exec.CommandContext(ctx,
		fmt.Sprintf("%s/bin/goimports", build.Default.GOPATH),
		"-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
