package app

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

var typicalwTmpl = `#!/bin/bash

set -e

TYPTMP={{.TypicalTmp}}
TYPGO=$TYPTMP/bin/typical-go

if ! [ -s $TYPGO ]; then
	echo "Build typical-go"
	go build -o $TYPGO github.com/typical-go/typical-go
fi

$TYPGO run \
	-project-pkg="{{.ProjectPkg}}" \
	-typical-build="{{.TypicalBuild}}" \
	-typical-tmp=$TYPTMP \
	$@
`

func cmdSetup() *cli.Command {
	return &cli.Command{
		Name:  "setup",
		Usage: "Setup typical-go",
		Flags: []cli.Flag{
			projectPkgFlag,
			typicalBuildFlag,
			typicalTmpFlag,
			&cli.StringFlag{Name: "gomod", Usage: "Iniate go.mod before setup if not empty"},
			&cli.BoolFlag{Name: "new", Usage: "Setup new project with standard layout and typical-build"},
		},
		Action: Setup,
	}
}

// Setup typical-go
func Setup(c *cli.Context) error {
	if gomod := c.String("gomod"); gomod != "" {
		if err := initGoMod(c.Context, gomod); err != nil {
			return err
		}
	}

	p, err := GetParam(c)
	if err != nil {
		return err
	}

	if c.Bool("new") {
		if err := newProject(p); err != nil {
			return err
		}
	}
	return createWrapper(p)
}

func initGoMod(ctx context.Context, pkg string) error {
	var stderr strings.Builder
	fmt.Fprintf(Stdout, "Initiate go.mod\n")
	if err := execkit.Run(ctx, &execkit.Command{
		Name:   "go",
		Args:   []string{"mod", "init", pkg},
		Stderr: &stderr,
	}); err != nil {
		return fmt.Errorf("%s: %s", err.Error(), stderr.String())
	}
	return nil
}

func createWrapper(p *Param) error {
	path := fmt.Sprintf("%s/typicalw", p.ProjectDir)
	fmt.Fprintf(Stdout, "Create wrapper '%s'\n", path)
	return common.ExecuteTmplToFile(path, typicalwTmpl, p)
}

func newProject(p *Param) error {
	projectName := filepath.Base(p.ProjectPkg)

	mkdirAll(p.TypicalTmp)
	mkdirAll("cmd/" + projectName)
	mkdirAll("internal/app")
	mkdirAll("app")

	// TODO: write main function
	// TODO: write simple app
	// TODO: write typical-build

	return nil
}

func mkdirAll(path string) {
	fmt.Fprintf(Stdout, "Mkdir %s\n", path)
	os.MkdirAll(path, 0777)
}
