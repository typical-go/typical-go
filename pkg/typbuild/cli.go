package typbuild

import (
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/wrapper"
	"github.com/urfave/cli/v2"
)

func createBuildToolCli(b *BuildTool, c *Context) *cli.App {

	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = "Build-Tool"
	app.Description = c.Description
	app.Version = c.Version

	app.Before = func(cli *cli.Context) (err error) {
		filename := "typical/precond_DO_NOT_EDIT.go"
		ctx := cli.Context

		c := &PreconditionContext{
			Precond: typtmpl.Precond{
				Imports: retrImports(c),
			},
			Logger: typlog.Logger{
				Name:  "PRECOND",
				Color: typlog.DefaultColor,
			},
			Core: c,
			Ctx:  ctx,
		}

		if err = b.Precondition(c); err != nil {
			return
		}

		if len(c.Lines) > 0 {
			if err = typtmpl.WriteFile(filename, 0777, c); err != nil {
				return
			}
			if err = buildkit.NewGoImports(wrapper.TypicalTmp, filename).Execute(ctx); err != nil {
				return
			}
		} else {
			c.Info("No precondition")
			os.Remove(filename)
		}

		return
	}
	app.Commands = b.Commands(c)

	return app
}

func retrImports(c *Context) []string {
	imports := []string{
		"github.com/typical-go/typical-go/pkg/typapp",
	}
	for _, dir := range c.AppDirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, importDef(c, dir))
		}
	}
	return imports
}

func importDef(c *Context, dir string) string {
	return fmt.Sprintf("%s/%s", wrapper.ProjectPkg, dir)
}
