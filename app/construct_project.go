package app

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/typical-go/typical-go/app/internal/tmpl"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/runn"
	"github.com/typical-go/typical-go/pkg/runn/stdrun"
	"github.com/typical-go/typical-go/pkg/typenv"
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

func constructProject(c *cli.Context) (err error) {
	pkg := c.Args().First()
	if pkg == "" {
		return cli.ShowCommandHelp(c, "new")
	}
	name := filepath.Base(pkg)
	if common.IsFileExist(name) {
		return fmt.Errorf("'%s' already exist", name)
	}
	return runn.Run(constructproj{
		Name:  name,
		Pkg:   pkg,
		blank: c.Bool("blank"),
		ctx:   c.Context,
	})
}

type constructproj struct {
	Name  string
	Pkg   string
	blank bool
	ctx   context.Context
}

func (i constructproj) Path(s string) string {
	return fmt.Sprintf("%s/%s", i.Name, s)
}

func (i constructproj) Run() (err error) {
	return runn.Run(
		i.appPackage,
		i.cmdPackage,
		i.projectDescriptor,
		i.ignoreFile,
		wrapper(i.Name),
		stdrun.NewGoFmt(i.ctx, "./..."),
		i.gomod,
	)
}

func (i constructproj) appPackage() error {
	stmts := []interface{}{
		stdrun.NewMkdir(i.Path("app")),
	}
	if !i.blank {
		stmts = append(stmts,
			stdrun.NewMkdir(i.Path("app/config")),
			stdrun.NewWriteString(i.Path("app/config/config.go"), tmpl.Config),
			stdrun.NewWriteTemplate(i.Path("app/app.go"), tmpl.App, i),
		)
	}
	return runn.Run(stmts...)
}

func (i constructproj) projectDescriptor() error {
	var writeStmt interface{}
	path := "typical/descriptor.go"
	if i.blank {
		writeStmt = stdrun.NewWriteTemplate(i.Path(path), tmpl.Context, i)
	} else {
		writeStmt = stdrun.NewWriteTemplate(i.Path(path), tmpl.ContextWithAppModule, i)
	}
	return runn.Run(
		stdrun.NewMkdir(i.Path("typical")),
		writeStmt,
	)
}

func (i constructproj) cmdPackage() error {
	appMainPath := fmt.Sprintf("%s/%s", typenv.Layout.Cmd, i.Name)
	buildtoolMainPath := fmt.Sprintf("%s/%s-%s", typenv.Layout.Cmd, i.Name, typenv.BuildTool)
	data := tmpl.MainData{ImportTypical: i.Pkg + "/typical"}
	return runn.Run(
		stdrun.NewMkdir(i.Path(typenv.Layout.Cmd)),
		stdrun.NewMkdir(i.Path(appMainPath)),
		stdrun.NewMkdir(i.Path(buildtoolMainPath)),
		stdrun.NewWriteTemplate(i.Path(appMainPath+"/main.go"), tmpl.MainAppSrc, data),
		stdrun.NewWriteTemplate(i.Path(buildtoolMainPath+"/main.go"), tmpl.MainBuildToolSrc, data),
	)
}

func (i constructproj) ignoreFile() error {
	return runn.Run(
		stdrun.NewWriteString(i.Path(".gitignore"), tmpl.Gitignore).WithPermission(0700),
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
	return runn.Run(
		stdrun.NewWriteTemplate(i.Path("go.mod"), tmpl.GoMod, data),
	)
}
