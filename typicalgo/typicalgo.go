package typicalgo

import (
	"errors"
	"go/build"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/buildkit"
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
			Flags: []cli.Flag{
				&cli.BoolFlag{Name: "blank", Usage: "Create blank new project"},
			},
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
			Action: func(c *cli.Context) (err error) {
				first := c.Args().First()
				if first == "" {
					return errors.New("Missing the first argument for temp-folder path")
				}
				modulePackage := c.Args().Get(1)
				if modulePackage == "" {
					modulePackage = defaultModulePackage()
				}
				return wrapMe(c.Context, &wrapContext{
					Descriptor:    d,
					TempFolder:    typcore.TempFolder(first),
					modulePackage: modulePackage,
				})
			},
		},
	}
	return app.Run(os.Args)
}

func defaultModulePackage() (pkg string) {
	var (
		gomod *buildkit.GoMod
		err   error
		root  string
	)
	if root, err = os.Getwd(); err != nil {
		log.Warn("Can't get default module package. Failed to get working directory")
		return ""
	}
	if gomod, err = buildkit.CreateGoMod(root + "/go.mod"); err != nil {
		// NOTE: go.mod is not exist. Check if the project sit in $GOPATH
		gopath := build.Default.GOPATH
		if strings.HasPrefix(root, gopath) {
			return root[len(gopath):]
		}

		log.Warn("Can't get default module package. `go.mod` is missing and the project not in $GOPATH")
		return ""
	}

	return gomod.ModulePackage
}
