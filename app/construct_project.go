package app

import (
	"fmt"
	"path/filepath"

	"github.com/typical-go/typical-go/app/internal/tmpl"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/filekit"
	"github.com/typical-go/typical-go/pkg/utility/golang"
	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/typical-go/typical-go/pkg/utility/runner"
	"github.com/urfave/cli/v2"
)

func cmdConstructProject() *cli.Command {
	return &cli.Command{
		Name:      "new",
		Usage:     "Construct New Project",
		UsageText: "app new [Package]",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "blank", Usage: "Create blank new project"},
		},
		Action: constructProject,
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
	return runn.Execute(constructproj{
		Name:  name,
		Pkg:   pkg,
		blank: ctx.Bool("blank"),
	})
}

type constructproj struct {
	Name  string
	Pkg   string
	blank bool
}

func (i constructproj) Path(s string) string {
	return fmt.Sprintf("%s/%s", i.Name, s)
}

func (i constructproj) Run() (err error) {
	return runn.Execute(
		i.appPackage,
		i.cmdPackage,
		i.typicalContext,
		i.ignoreFile,
		wrapperRunner(i.Name),
		runner.NewGoFmt("./..."),
		i.gomod,
	)
}

func (i constructproj) appPackage() error {
	stmts := []interface{}{
		runner.NewMkdir(i.Path("app")),
	}
	if !i.blank {
		stmts = append(stmts,
			runner.NewMkdir(i.Path("app/config")),
			runner.NewWriteString(i.Path("app/config/config.go"), tmpl.Config),
			runner.NewWriteTemplate(i.Path("app/app.go"), tmpl.App, i),
			runner.NewWriteTemplate(i.Path("app/app_test.go"), tmpl.AppTest, i),
		)
	}
	return runn.Execute(stmts...)
}

func (i constructproj) typicalContext() error {
	var writeStmt interface{}
	path := "typical/context.go"
	if i.blank {
		writeStmt = runner.NewWriteTemplate(i.Path(path), tmpl.Context, i)
	} else {
		writeStmt = runner.NewWriteTemplate(i.Path(path), tmpl.ContextWithAppModule, i)
	}
	return runn.Execute(
		runner.NewMkdir(i.Path("typical")),
		writeStmt,
	)
}

func (i constructproj) cmdPackage() error {
	appMainPath := fmt.Sprintf("%s/%s", typenv.Layout.Cmd, i.Name)
	buildtoolMainPath := fmt.Sprintf("%s/%s-%s", typenv.Layout.Cmd, i.Name, typenv.BuildTool)
	return runn.Execute(
		runner.NewMkdir(i.Path(typenv.Layout.Cmd)),
		runner.NewMkdir(i.Path(appMainPath)),
		runner.NewMkdir(i.Path(buildtoolMainPath)),
		runner.NewWriteSource(i.Path(appMainPath+"/main.go"), i.appMainSrc()),
		runner.NewWriteSource(i.Path(buildtoolMainPath+"/main.go"), i.buildtoolMainSrc()),
	)
}

func (i constructproj) appMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typapp")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Append("typapp.Run(typical.Context)")
	return
}

func (i constructproj) buildtoolMainSrc() (src *golang.MainSource) {
	src = golang.NewMainSource()
	src.Imports.Add("", "github.com/typical-go/typical-go/pkg/typbuildtool")
	src.Imports.Add("", i.Pkg+"/typical")
	src.Append("typbuildtool.Run(typical.Context)")
	return
}

func (i constructproj) ignoreFile() error {
	return runn.Execute(
		runner.NewWriteString(i.Path(".gitignore"), tmpl.Gitignore).WithPermission(0700),
	)
}

func (i constructproj) gomod() (err error) {
	data := struct {
		Pkg            string
		TypicalVersion string
	}{
		Pkg:            i.Pkg,
		TypicalVersion: Version,
	}
	return runn.Execute(
		runner.NewWriteTemplate(i.Path("go.mod"), tmpl.GoMod, data),
	)
}
