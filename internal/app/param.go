package app

import (
	"context"
	"errors"
	"path/filepath"
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
		AppName  string
		SetupTarget  string
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
	typicalBuildFlag = &cli.StringFlag{
		Name:  TypicalBuildParam,
		Usage: "Typical-Build source code location",
		Value: DefaultTypicalBuild,
	}
	projectPkgFlag = &cli.StringFlag{
		Name:  ProjectPkgParam,
		Usage: "Project package name. Same with module package in go.mod by default",
	}
)

// GetParam get param
func GetParam(c *cli.Context) (*Param, error) {
	projectPkg := c.String(ProjectPkgParam)
	setupTarget := ""
	if projectPkg == "" {
		var err error
		projectPkg, err = retrieveProjPkg(c.Context)
		if err != nil {
			return nil, err
		}
		setupTarget = "."
	} else {
		setupTarget = filepath.Base(projectPkg)
	}

	return &Param{
		TypicalBuild: c.String(TypicalBuildParam),
		TypicalTmp:   c.String(TypicalTmpParam),
		ProjectPkg:   projectPkg,
		AppName:  filepath.Base(projectPkg),
		SetupTarget:  setupTarget,
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
