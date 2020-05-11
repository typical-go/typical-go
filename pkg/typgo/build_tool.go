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

// BuildTool detail
type BuildTool struct {
	*Descriptor

	AppDirs  []string
	AppFiles []string
}

// ActionFunc to return related action func
func (c *BuildTool) ActionFunc(name string, fn CliFunc) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		return fn(&CliContext{
			Logger: typlog.Logger{
				Name: strings.ToUpper(name),
			},
			Cli:       cli,
			BuildTool: c,
		})
	}
}

func createBuildTool(d *Descriptor) *BuildTool {
	appDirs, appFiles := WalkLayout(d.Layouts)

	return &BuildTool{
		Descriptor: d,
		AppDirs:    appDirs,
		AppFiles:   appFiles,
	}
}

func launchBuildTool(d *Descriptor) error {

	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "Build-Tool"
	app.Description = d.Description
	app.Version = d.Version

	c := createBuildTool(d)

	app.Before = func(cli *cli.Context) (err error) {

		os.Remove(typvar.PrecondFile)

		ctx := cli.Context

		c := &PrecondContext{
			Precond: typtmpl.Precond{
				Imports: retrImports(c),
			},
			Logger:    typlog.Logger{Name: "PRECOND", Color: typlog.DefaultColor},
			BuildTool: c,
			Ctx:       ctx,
		}

		if err = c.Precondition(c); err != nil {
			return
		}

		if len(c.Lines) > 0 {
			if err = typtmpl.WriteFile(typvar.PrecondFile, 0777, c); err != nil {
				return
			}
			if err = buildkit.NewGoImports(typvar.TypicalTmp, typvar.PrecondFile).Execute(ctx); err != nil {
				return
			}
		} else {
			c.Info("No precondition")
			os.Remove(typvar.PrecondFile)
		}

		return
	}
	app.Commands = c.Commands(c)

	return app.Run(os.Args)
}

func retrImports(c *BuildTool) []string {
	imports := []string{
		"github.com/typical-go/typical-go/pkg/typgo",
	}
	for _, dir := range c.AppDirs {
		if !strings.Contains(dir, "internal") {
			imports = append(imports, fmt.Sprintf("%s/%s", typvar.ProjectPkg, dir))
		}
	}
	return imports
}
