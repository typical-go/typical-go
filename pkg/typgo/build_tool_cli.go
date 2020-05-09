package typgo

import (
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

var (
	precondFile = "typical/precond_DO_NOT_EDIT.go"
)

func createBuildToolCli(b *BuildTool, c *Context) *cli.App {

	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = "Build-Tool"
	app.Description = c.Description
	app.Version = c.Version

	app.Before = func(cli *cli.Context) (err error) {

		os.Remove(precondFile)

		ctx := cli.Context

		c := &PrecondContext{
			Precond: typtmpl.Precond{
				Imports: retrImports(c),
			},
			Logger: typlog.Logger{
				Name:  "PRECOND",
				Color: typlog.DefaultColor,
			},
			Context: c,
			Ctx:     ctx,
		}

		if err = b.Precondition(c); err != nil {
			return
		}

		if len(c.Lines) > 0 {
			if err = typtmpl.WriteFile(precondFile, 0777, c); err != nil {
				return
			}
			if err = buildkit.NewGoImports(typvar.TypicalTmp, precondFile).Execute(ctx); err != nil {
				return
			}
		} else {
			c.Info("No precondition")
			os.Remove(precondFile)
		}

		return
	}
	app.Commands = b.Commands(c)

	return app
}

func retrImports(c *Context) []string {
	imports := []string{
		"github.com/typical-go/typical-go/pkg/typgo",
	}
	for _, dir := range c.AppDirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, importDef(c, dir))
		}
	}
	return imports
}

func importDef(c *Context, dir string) string {
	return fmt.Sprintf("%s/%s", typvar.ProjectPkg, dir)
}
