package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
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
			Usage: "wrap the project and run the build-tool",
			Flags: []cli.Flag{
				&cli.StringFlag{Name: typicalTmpParam, Value: ".typical-tmp"},
				&cli.StringFlag{Name: srcParam, Value: "tools/typical-build"},
				&cli.StringFlag{Name: projPkgParam},
			},
			Action: func(c *cli.Context) (err error) {
				typicalTmp := c.String(typicalTmpParam)
				projectPkg := c.String(projPkgParam)
				src := c.String(srcParam)

				if projectPkg == "" {
					if projectPkg, err = retrieveProjPkg(c.Context); err != nil {
						return err
					}
				}

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

func retrieveProjPkg(ctx context.Context) (string, error) {
	var stderr strings.Builder
	var stdout strings.Builder
	cmd := execkit.Command{
		Name:   "go",
		Args:   []string{"list", "-m"},
		Stdout: &stdout,
		Stderr: &stderr,
	}
	if err := cmd.Run(ctx); err != nil {
		return "", errors.New(stderr.String())
	}
	return strings.TrimSpace(stdout.String()), nil
}
