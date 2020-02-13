package app

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/typical-go/typical-go/app/internal/tmpl"
	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/runn"
	"github.com/typical-go/typical-go/pkg/runn/stdrun"
	"github.com/typical-go/typical-go/pkg/typcore"
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
		TemplateData: tmpl.TemplateData{
			Name: name,
			Pkg:  pkg,
		},
		blank: c.Bool("blank"),
		ctx:   c.Context,
	})
}

type constructproj struct {
	tmpl.TemplateData
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
		i.descriptor,
		i.ignoreFile,
		wrapper(i.Name, i.Pkg),
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
			stdrun.NewWriteTemplate(i.Path("app/app.go"), tmpl.App, i.TemplateData),
		)
	}
	return runn.Run(stmts...)
}

func (i constructproj) descriptor() error {
	var writeStmt interface{}
	path := "typical/descriptor.go"
	if i.blank {
		writeStmt = stdrun.NewWriteTemplate(i.Path(path), tmpl.Context, i.TemplateData)
	} else {
		writeStmt = stdrun.NewWriteTemplate(i.Path(path), tmpl.ContextWithAppModule, i.TemplateData)
	}
	return runn.Run(
		stdrun.NewMkdir(i.Path("typical")),
		writeStmt,
	)
}

func (i constructproj) cmdPackage() error {
	appMainPath := fmt.Sprintf("%s/%s", typcore.DefaultLayout.Cmd, i.Name)
	data := tmpl.MainSrcData{
		ImportTypical: i.Pkg + "/typical",
	}
	return runn.Run(
		stdrun.NewMkdir(i.Path(typcore.DefaultLayout.Cmd)),
		stdrun.NewMkdir(i.Path(appMainPath)),
		stdrun.NewWriteTemplate(i.Path(appMainPath+"/main.go"), tmpl.MainSrcApp, data),
	)
}

func (i constructproj) ignoreFile() error {
	return runn.Run(
		stdrun.NewWriteString(i.Path(".gitignore"), tmpl.Gitignore).WithPermission(0700),
	)
}

func (i constructproj) gomod() (err error) {
	return runn.Run(
		stdrun.NewWriteTemplate(i.Path("go.mod"), tmpl.GoMod, tmpl.GoModData{
			Pkg:            i.Pkg,
			TypicalVersion: Version,
		}),
	)
}
