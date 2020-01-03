package typbuildtool

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/golang"
	"github.com/typical-go/typical-go/pkg/runn/stdrun"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func (t buildtool) prebuild(c *cli.Context) (err error) {
	log.Info("Walk the project")
	var (
		autowires Autowires
		ctx       = c.Context
	)
	walker := walker.New(t.filenames)
	walker.AddFuncDeclListener(&autowires)
	if err = walker.Walk(); err != nil {
		return
	}
	log.Info("Generate constructors")
	if err = t.generateConstructor(ctx, typenv.AppMainPath+"/constructor.go", autowires); err != nil {
		return
	}
	return
}

func (t buildtool) generateConstructor(ctx context.Context, target string, constructors []string) (err error) {
	defer common.ElapsedTimeFn("Generate constructor")()
	src := golang.NewSource("main")
	if len(constructors) < 1 {
		return
	}
	imports := make(map[string]struct{})
	imports[t.Package+"/typical"] = struct{}{}
	for _, dir := range t.dirs {
		imports[t.Package+"/"+dir] = struct{}{}
	}
	for _, constructor := range constructors {
		src.Init.Append(fmt.Sprintf("typical.Descriptor.Constructors.Append(%s)", constructor))
	}
	for key := range imports {
		src.Imports.Add("", key)
	}
	if err = stdrun.NewWriteSource(target, src).Run(); err != nil {
		return
	}
	cmd := exec.CommandContext(ctx,
		fmt.Sprintf("%s/bin/goimports", build.Default.GOPATH),
		"-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
