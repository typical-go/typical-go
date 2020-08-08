package app

import (
	"context"
	"errors"
	"strings"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

type (
	// Param for typical-go
	Param struct {
		TypicalBuild string
		TypicalTmp   string
		ProjectPkg   string
	}
)

var (
	// TypicalTmpParam typical-tmp param
	TypicalTmpParam = "typical-tmp"
	// DefaultTypicalTmp typical-tmp default value
	DefaultTypicalTmp = ".typical-tmp"
	// TypicalBuildParam typical-build param
	TypicalBuildParam = "typical-build"
	// DefaultTypicalBuild typical-build default value
	DefaultTypicalBuild = "tools/typical-build"
	// ProjectPkgParam project-pkg param
	ProjectPkgParam = "project-pkg"
	typicalTmpFlag  = &cli.StringFlag{
		Name:  TypicalTmpParam,
		Usage: "Temporary directory location to save builds-related files",
		Value: DefaultTypicalTmp,
	}
	projPkgFlag = &cli.StringFlag{
		Name:  ProjectPkgParam,
		Usage: "Project package name. Same with module package in go.mod by default",
	}
	srcFlag = &cli.StringFlag{
		Name:  TypicalBuildParam,
		Usage: "Typical-Build source code location",
		Value: DefaultTypicalBuild,
	}
)

// GetParam get param
func GetParam(c *cli.Context) (*Param, error) {
	projPkg := c.String(ProjectPkgParam)
	if projPkg == "" {
		var err error
		projPkg, err = retrieveProjPkg(c.Context)
		if err != nil {
			return nil, err
		}
	}

	return &Param{
		TypicalBuild: c.String(TypicalBuildParam),
		TypicalTmp:   c.String(TypicalTmpParam),
		ProjectPkg:   projPkg,
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
		return "", errors.New(err.Error() + ": " + stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}
