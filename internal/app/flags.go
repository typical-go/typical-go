package app

import (
	"context"
	"errors"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
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

func getSrc(c *cli.Context) string {
	return c.String(srcParam)
}

func getTypicalTmp(c *cli.Context) string {
	return c.String(typicalTmpParam)
}

func getProjectPkg(c *cli.Context) (s string, err error) {
	projPkg := c.String(projPkgParam)
	if projPkg == "" {
		if projPkg, err = retrieveProjPkg(c.Context); err != nil {
			return "", err
		}
	}
	return projPkg, nil
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
