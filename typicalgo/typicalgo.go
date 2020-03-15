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

// Run the typical-go
func (t *TypicalGo) Run(d *typcore.Descriptor) (err error) {
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version

	app.Commands = []*cli.Command{
		{
			Name: "wrap-me",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "tmp", Required: true},
				&cli.StringFlag{Name: "project-package"},
			},
			Action: func(cliCtx *cli.Context) (err error) {

				return wrapMe(&wrapContext{
					Context:        typcore.CreateContext(d),
					Cli:            cliCtx,
					tmp:            cliCtx.String("tmp"),
					projectPackage: cliCtx.String("project-package"),
				})
			},
		},
	}
	return app.Run(os.Args)
}
