package typicalgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// TypicalGo is app of typical-go
type TypicalGo struct{}

// New of Typical-Go
func New() *TypicalGo {
	return &TypicalGo{}
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
			Name: "wrap",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "tmp-folder", Required: true},
				&cli.StringFlag{Name: "project-package", Usage: "To override generated ProjectPackage in context"},
			},
			Action: func(cliCtx *cli.Context) (err error) {

				return d.Wrap(&typcore.WrapContext{
					Descriptor:     d,
					Ctx:            cliCtx.Context,
					TmpFolder:      cliCtx.String("tmp-folder"),
					ProjectPackage: cliCtx.String("project-package"),
				})
			},
		},
	}
	return app.Run(os.Args)
}
