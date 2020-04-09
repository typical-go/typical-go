package wrapper

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

var (
	_ typcore.App = (*App)(nil)
)

// App of wrapper
type App struct {
}

// New instance of TypicalGo
func New() *App {
	return &App{}
}

// RunApp to run the typical-go
func (t *App) RunApp(d *typcore.Descriptor) (err error) {
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version

	app.Commands = []*cli.Command{
		{
			Name:  "wrap",
			Usage: "wrap the project with its build-tool",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "typical-tmp", Value: ".typical-tmp"},
				&cli.StringFlag{Name: "descriptor-folder", Value: "typical"},
				&cli.StringFlag{Name: "checksum-file", Value: "checksum"},
				&cli.StringFlag{Name: "project-pkg", Usage: "To override generated ProjectPackage in context"},
			},
			Action: func(cliCtx *cli.Context) (err error) {
				return Wrap(&Context{
					Descriptor:       d,
					Ctx:              cliCtx.Context,
					TypicalTmp:       cliCtx.String("typical-tmp"),
					ProjectPkg:       cliCtx.String("project-pkg"),
					DescriptorFolder: cliCtx.String("descriptor-folder"),
					ChecksumFile:     cliCtx.String("checksum-file"),
				})
			},
		},
	}
	return app.Run(os.Args)
}

// AppSources is application source for typical-go
func (t *App) AppSources() []string {
	return []string{"wrapper"}
}
