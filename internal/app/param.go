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
		ProjectDir   string
		ProjectName  string
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
	// ProjectDirParam project-dir param
	ProjectDirParam = "project-dir"
	// DefaultProjectDir project-dir default value
	DefaultProjectDir = "."
	typicalTmpFlag    = &cli.StringFlag{
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
	projectDirFlag = &cli.StringFlag{
		Name:  ProjectDirParam,
		Usage: "Project directory location",
		Value: ".",
	}
)

// GetParam get param
func GetParam(c *cli.Context) (*Param, error) {
	projectDir := c.String(ProjectDirParam)
	projectPkg := c.String(ProjectPkgParam)
	if projectPkg == "" {
		var err error
		projectPkg, err = retrieveProjPkg(c.Context, projectDir)
		if err != nil {
			return nil, err
		}
	}

	return &Param{
		TypicalBuild: c.String(TypicalBuildParam),
		TypicalTmp:   c.String(TypicalTmpParam),
		ProjectPkg:   projectPkg,
		ProjectDir:   projectDir,
		ProjectName:  filepath.Base(projectPkg),
	}, nil
}

func retrieveProjPkg(ctx context.Context, projectDir string) (string, error) {
	var stdout strings.Builder
	var stderr strings.Builder
	if err := execkit.Run(ctx, &execkit.Command{
		Name:   "go",
		Args:   []string{"list", "-m"},
		Stdout: &stdout,
		Stderr: &stderr,
		Dir:    projectDir,
	}); err != nil {
		return "", errors.New(err.Error() + ": " + stderr.String())
	}

	return strings.TrimSpace(stdout.String()), nil
}
