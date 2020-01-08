package typbuildtool

import (
	"context"
	"fmt"
	"go/build"
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/typcore"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/golang"
	"github.com/typical-go/typical-go/pkg/runn/stdrun"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
	"github.com/typical-go/typical-go/pkg/typenv"
)

func prebuild(ctx context.Context, d *typcore.ProjectDescriptor) (err error) {
	var (
		autowires Autowires
		filenames []string
		dirs      []string
	)
	if dirs, filenames, err = projectFiles(typenv.Layout.App); err != nil {
		log.Fatal(err.Error())
	}
	log.Info("Walk the project")
	walker := walker.New(filenames).
		AddDeclListener(&autowires)
	if err = walker.Walk(); err != nil {
		return
	}
	// TODO: generate imports
	log.Info("Generate constructors")
	if err = generateConstructor(ctx, d, typenv.AppMainPath+"/constructor.go", autowires, dirs); err != nil {
		return
	}
	return
}

func generateConstructor(ctx context.Context, d *typcore.ProjectDescriptor, target string, constructors, dirs []string) (err error) {
	defer common.ElapsedTimeFn("Generate constructor")()
	src := golang.NewSource("main")
	if len(constructors) < 1 {
		return
	}
	imports := make(map[string]struct{})
	imports[d.Package+"/typical"] = struct{}{}
	for _, dir := range dirs {
		imports[d.Package+"/"+dir] = struct{}{}
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
