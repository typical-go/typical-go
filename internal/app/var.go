package app

import (
	"context"
	"errors"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

const (
	typicalTmpParam    = "typical-tmp"
	projPkgParam       = "project-pkg"
	srcParam           = "src"
	createWrapperParam = "create:wrapper"
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
