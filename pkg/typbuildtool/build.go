package typbuildtool

import (
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
	"github.com/typical-go/typical-go/pkg/utility/debugkit"
	"github.com/typical-go/typical-go/pkg/utility/filekit"
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
	t.generateConstructor(typenv.AppMainPath+"/constructor.go", autowires)
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
	for _, constructor := range constructors {
		dotIndex := strings.Index(constructor, ".")
		if dotIndex >= 0 {
			pkg := constructor[:dotIndex]
			imports[t.Package+"/"+pkg] = struct{}{}
		}
		src.Init.Append(fmt.Sprintf("typical.Context.Constructors.Append(%s)", constructor))
	}
	for key := range imports {
		src.Imports.Add("", key)
	}
	if err = filekit.Write(target, src); err != nil {
		return
	}
	return goimports(target)
}

func goimports(filename string) error {
	// TODO: change ot gofmt
	cmd := exec.Command(fmt.Sprintf("%s/bin/goimports", build.Default.GOPATH),
		"-w", filename)
	return cmd.Run()
}
