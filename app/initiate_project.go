package app

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/typical-go/typical-go/pkg/utility/golang"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/app/common"
	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/urfave/cli"
)

func initiateProject(ctx *cli.Context) error {
	pkg := "github.com/typical-go/sample"
	name := name(pkg)
	log.Infof("Remove: %s", name)
	os.RemoveAll(name)
	log.Infof("Init Project: %s", pkg)
	return runn.Execute(initproject{
		Name: name,
		Pkg:  pkg,
	})
}

func name(pkg string) string {
	// TODO: handle window path format
	chunks := strings.Split(pkg, "/")
	return chunks[len(chunks)-1]
}

type initproject struct {
	Name string
	Pkg  string
}

func (i initproject) Path(s string) string {
	return fmt.Sprintf("%s/%s", i.Name, s)
}

func (i initproject) Run() (err error) {
	return runn.Execute(
		i.generateAppPackage,
		i.generateCmdPackage,
		i.generateTypicalContext,
		i.generateIgnoreFile,
		i.generateTypicalWrapper,
		i.initGoModule,
	)
}

func (i initproject) generateAppPackage() error {
	log.Info("Generate App Package")
	return runn.Execute(
		common.Mkdir{Path: i.Path("app")},
	)
}

func (i initproject) generateCmdPackage() error {
	log.Info("Generate Cmd Package")

	return runn.Execute(
		common.Mkdir{Path: i.Path("cmd")},
		common.Mkdir{Path: i.Path("cmd/app")},
		common.Mkdir{Path: i.Path("cmd/pre-builder")},
		common.Mkdir{Path: i.Path("cmd/build-tool")},
		common.WriteSource{Target: i.Path("cmd/app/main.go"), Source: i.appMainSrc()},
		common.WriteSource{Target: i.Path("cmd/pre-builder/main.go"), Source: i.prebuilderMainSrc()},
		common.WriteSource{Target: i.Path("cmd/build-tool/main.go"), Source: i.buildtoolMainSrc()},
	)
}

func (i initproject) appMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typapp")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Imports.Add("_", i.Pkg+"/internal/dependency")
	src.Append("typapp.Run(typical.Context)")
	return
}

func (i initproject) prebuilderMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typprebuilder")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Append("typprebuilder.Run(typical.Context)")
	return
}

func (i initproject) buildtoolMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typbuildtool")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Imports.Add("_", i.Pkg+"/internal/dependency")
	src.Append("typbuildtool.Run(typical.Context)")
	return
}

func (i initproject) generateTypicalContext() error {
	log.Info("Generate Typical Context")
	template := `package typical

import (
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typrls"
)

// Context of Project
var Context = &typctx.Context{
	Name:    "{{.Name}}",
	Version: "0.0.1",
	Package: "{{.Pkg}}",
	Releaser: typrls.Releaser{
		Targets: []typrls.Target{"linux/amd64", "darwin/amd64"},
	},
}
`

	return runn.Execute(
		common.Mkdir{Path: i.Path("typical")},
		common.WriteTemplate{
			Target:   i.Path("typical/context.go"),
			Template: template,
			Data:     i,
		},
	)
}

func (i initproject) generateIgnoreFile() error {
	log.Info("Generate Ignore File")
	return runn.Execute()
}

func (i initproject) generateTypicalWrapper() error {
	log.Info("Generate Typical Wrapper")
	return runn.Execute()
}

func (i initproject) initGoModule() (err error) {
	cmd := exec.Command("go", "mod", "init", i.Pkg)
	cmd.Dir = i.Name
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
