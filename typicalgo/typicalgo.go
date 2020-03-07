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
			Name:      "new",
			Usage:     "Construct New Project",
			UsageText: "app new [Package]",
			Action: func(c *cli.Context) (err error) {
				pkg := c.Args().First()
				if pkg == "" {
					return cli.ShowCommandHelp(c, "new")
				}
				return constructProject(c.Context, pkg)
			},
		},
		{
			Name: "wrap-me",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: "tmp", Required: true},
				&cli.StringFlag{Name: "module-package"},
				&cli.StringFlag{Name: "name"},
			},
			Action: func(c *cli.Context) (err error) {
				var (
					root          string
					tmp           = c.String("tmp")
					modulePackage = c.String("module-package")
				)

				if root, err = os.Getwd(); err != nil {
					return err
				}

				if modulePackage == "" {
					modulePackage = typcore.RetrieveModulePackage(root)
				}

				return wrapMe(c.Context, &wrapContext{
					Descriptor:    d,
					TempFolder:    typcore.TempFolder(tmp),
					modulePackage: modulePackage,
				})
			},
		},
	}
	return app.Run(os.Args)
}
