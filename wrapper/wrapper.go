package wrapper

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

var (
	_ typcore.App = (*Wrapper)(nil)
)

// Wrapper is app of typical-go
type Wrapper struct {
}

// New instance of TypicalGo
func New() *Wrapper {
	return &Wrapper{}
}

// RunApp to run the typical-go
func (t *Wrapper) RunApp(d *typcore.Descriptor) (err error) {
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
				&cli.StringFlag{Name: "project-package", Usage: "To override generated ProjectPackage in context"},
			},
			Action: func(cliCtx *cli.Context) (err error) {
				return Wrap(&Context{
					Descriptor:     d,
					Ctx:            cliCtx.Context,
					TypicalTmp:     cliCtx.String("typical-tmp"),
					ProjectPackage: cliCtx.String("project-package"),
				})
			},
		},
	}
	return app.Run(os.Args)
}

// AppSources is application source for typical-go
func (t *Wrapper) AppSources() []string {
	return []string{"wrapper"}
}
