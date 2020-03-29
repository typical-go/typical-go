package app

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typwrap"
	"github.com/urfave/cli/v2"
)

var (
	_ typcore.App = (*TypicalGo)(nil)
)

// TypicalGo is app of typical-go
type TypicalGo struct {
	wrapper *typwrap.Wrapper
}

// New instance of TypicalGo
func New() *TypicalGo {
	return &TypicalGo{
		wrapper: typwrap.New(),
	}
}

// RunApp to run the typical-go
func (t *TypicalGo) RunApp(d *typcore.Descriptor) (err error) {
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
				return t.Wrap(&typwrap.Context{
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

// Wrap the project
func (t *TypicalGo) Wrap(c *typwrap.Context) error {
	return t.wrapper.Wrap(c)
}

// AppSources is application source for typical-go
func (t *TypicalGo) AppSources() []string {
	return []string{"app"}
}
