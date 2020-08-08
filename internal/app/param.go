package app

import (
	"context"
	"errors"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

type (
	param struct {
		Src        string
		TypicalTmp string
		ProjectPkg string
	}
)

var (
	typicalTmpParam = "typical-tmp"
	projPkgParam    = "project-pkg"
	srcParam        = "src"

	typicalTmpFlag = &cli.StringFlag{
		Name:  typicalTmpParam,
		Usage: "Temporary directory location to save builds-related files",
		Value: ".typical-tmp",
	}

	projPkgFlag = &cli.StringFlag{
		Name:  projPkgParam,
		Usage: "Project package name. Same with module package in go.mod by default",
	}

	srcFlag = &cli.StringFlag{
		Name:  srcParam,
		Usage: "Build-tool source",
		Value: "tools/typical-build",
	}
)

func getParam(c *cli.Context) (*param, error) {
	projPkg := c.String(projPkgParam)
	if projPkg == "" {
		var err error
		projPkg, err = retrieveProjPkg(c.Context)
		if err != nil {
			return nil, err
		}
	}

	return &param{
		Src:        c.String(srcParam),
		TypicalTmp: c.String(typicalTmpParam),
		ProjectPkg: projPkg,
	}, nil
}

func retrieveProjPkg(ctx context.Context) (string, error) {
	var stdout strings.Builder
	var stderr strings.Builder
	if err := execkit.Run(ctx, &execkit.Command{
		Name:   "go",
		Args:   []string{"list", "-m"},
		Stdout: &stdout,
		Stderr: &stderr,
	}); err != nil {
		return "", errors.New(stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}
