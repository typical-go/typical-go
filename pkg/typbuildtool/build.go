package typbuildtool

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"

	"github.com/typical-go/typical-go/pkg/utility/runner"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/debugkit"
	"github.com/typical-go/typical-go/pkg/utility/golang"

	"github.com/urfave/cli/v2"
)

func (t buildtool) cmdBuild() *cli.Command {
	return &cli.Command{
		Name:    "build",
		Aliases: []string{"b"},
		Usage:   "Build the binary",
		Action:  t.buildBinary,
	}
}

func (t buildtool) buildBinary(ctx *cli.Context) (err error) {
	log.Info("Walk the project")
	var autowires Autowires
	walker := walker.New(t.filenames)
	walker.AddFuncDeclListener(&autowires)
	if err = walker.Walk(); err != nil {
		return
	}
	log.Info("Generate constructors")
	if err = t.generateConstructor(typenv.AppMainPath+"/constructor.go", autowires); err != nil {
		return
	}
	log.Info("Build the application")
	cmd := exec.Command("go", "build",
		"-o", typenv.AppBin,
		"./"+typenv.AppMainPath,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (t buildtool) generateConstructor(target string, constructors []string) (err error) {
	defer debugkit.ElapsedTime("Generate constructor")()
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
	if err = runner.NewWriteSource(target, src).Run(); err != nil {
		return
	}
	cmd := exec.Command(fmt.Sprintf("%s/bin/goimports", build.Default.GOPATH),
		"-w", target)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
