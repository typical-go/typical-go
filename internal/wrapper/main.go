package wrapper

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/urfave/cli/v2"
)

// Main function to run the typical-go
func Main() (err error) {

	app := cli.NewApp()
	app.Name = typapp.Name
	app.Usage = ""       // NOTE: intentionally blank
	app.Description = "" // NOTE: intentionally blank
	app.Version = typapp.Version

	app.Commands = []*cli.Command{
		{
			Name:  "wrap",
			Usage: "wrap the project with its build-tool",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: typicalTmpParam, Value: ".typical-tmp"},
				&cli.StringFlag{Name: srcParam, Value: "tools/typical-build"},
				&cli.StringFlag{Name: projPkgParam},
			},
			Action: func(c *cli.Context) error {
				typicalTmp := c.String(typicalTmpParam)
				projectPkg := c.String(projPkgParam)
				src := c.String(srcParam)

				return wrap(&wrapContext{
					Context:      c.Context,
					args:         c.Args().Slice(),
					typicalTmp:   typicalTmp,
					projectPkg:   projectPkg,
					src:          src,
					chksumTarget: fmt.Sprintf("%s/checksum", typicalTmp),
					bin:          fmt.Sprintf("%s/bin/%s", typicalTmp, filepath.Base(src)),
				})
			},
		},
	}

	return app.Run(os.Args)
}
