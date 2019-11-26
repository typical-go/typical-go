package app

import (
	"fmt"
	"os/exec"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/utility/filekit"
	"github.com/typical-go/typical-go/pkg/utility/golang"
	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/typical-go/typical-go/pkg/utility/runner"
	"github.com/urfave/cli"
)

func cmdConstructProject() cli.Command {
	return cli.Command{
		Name:      "new",
		Usage:     "Construct New Project",
		UsageText: "app new [Package]",
		Action:    constructProject,
	}
}

func constructProject(ctx *cli.Context) (err error) {
	pkg := ctx.Args().First()
	if pkg == "" {
		return cli.ShowCommandHelp(ctx, "new")
	}
	name := filepath.Base(pkg)
	if filekit.IsExist(name) {
		return fmt.Errorf("'%s' already exist", name)
	}
	return runn.Execute(
		constructproj{Name: name, Pkg: pkg},
	)
}

type constructproj struct {
	Name string
	Pkg  string
}

func (i constructproj) Path(s string) string {
	return fmt.Sprintf("%s/%s", i.Name, s)
}

func (i constructproj) Run() (err error) {
	return runn.Execute(
		i.appPackage,
		i.cmdPackage,
		i.dependencyPackage,
		i.typicalContext,
		i.ignoreFile,
		i.gomod,
		i.gofmt,
		wrapperRunner(i.Name),
	)
}

func (i constructproj) appPackage() error {
	return runn.Execute(
		runner.Mkdir{Path: i.Path("app")},
		runner.WriteString{
			Target:     i.Path("app/app.go"),
			Content:    appSrc,
			Permission: 0644,
		},
		runner.WriteTemplate{
			Target:   i.Path("app/app_test.go"),
			Template: appSrcTest,
			Data:     i,
		},
	)
}

func (i constructproj) cmdPackage() error {
	return runn.Execute(
		runner.Mkdir{Path: i.Path("cmd")},
		runner.Mkdir{Path: i.Path("cmd/app")},
		runner.Mkdir{Path: i.Path("cmd/pre-builder")},
		runner.Mkdir{Path: i.Path("cmd/build-tool")},
		runner.WriteSource{Target: i.Path("cmd/app/main.go"), Source: i.appMainSrc()},
		runner.WriteSource{Target: i.Path("cmd/pre-builder/main.go"), Source: i.prebuilderMainSrc()},
		runner.WriteSource{Target: i.Path("cmd/build-tool/main.go"), Source: i.buildtoolMainSrc()},
	)
}

func (i constructproj) appMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typapp")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Imports.Add("_", i.Pkg+"/internal/dependency")
	src.Append("typapp.Run(typical.Context)")
	return
}

func (i constructproj) prebuilderMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typprebuilder")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Append("typprebuilder.Run(typical.Context)")
	return
}

func (i constructproj) buildtoolMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typbuildtool")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Imports.Add("_", i.Pkg+"/internal/dependency")
	src.Append("typbuildtool.Run(typical.Context)")
	return
}

func (i constructproj) typicalContext() error {
	return runn.Execute(
		runner.Mkdir{Path: i.Path("typical")},
		runner.WriteTemplate{
			Target:   i.Path("typical/context.go"),
			Template: ctxSrc,
			Data:     i,
		},
	)
}

func (i constructproj) ignoreFile() error {
	return runn.Execute(
		runner.WriteString{
			Target:     i.Path(".gitignore"),
			Permission: 0700,
			Content:    gitignore,
		},
	)
}

func (i constructproj) dependencyPackage() error {
	return runn.Execute(
		runner.Mkdir{Path: i.Path("internal/dependency")},
		runner.WriteString{
			Target:     i.Path("internal/dependency/constructors.go"),
			Permission: 0644,
			Content:    "package dependency",
		},
	)
}

func (i constructproj) gomod() (err error) {
	return runn.Execute(
		runner.WriteTemplate{
			Target:   i.Path("go.mod"),
			Template: gomod,
			Data: struct {
				Pkg            string
				TypicalVersion string
			}{
				Pkg:            i.Pkg,
				TypicalVersion: Version,
			},
		},
	)
}

func (i constructproj) gofmt() (err error) {
	cmd := exec.Command("go", "fmt", "./...")
	cmd.Dir = i.Name
	return cmd.Run()
}
