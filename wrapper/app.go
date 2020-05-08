package wrapper

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/urfave/cli/v2"
)

// Main function to run the typical-go
func Main(d *typcore.Descriptor) (err error) {
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
					Descriptor: d,
					Logger: typlog.Logger{
						Name: "WRAPPER",
					},
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
