package app

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

var typicalw = "typicalw"

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
			srcFlag,
			projPkgFlag,
			typicalTmpFlag,
			&cli.StringFlag{Name: "gomod", Usage: "Iniate go.mod before setup if not empty"},
			&cli.StringFlag{Name: "new", Usage: "Setup new project with standard layout and typical-build"},
		},
		Action: setup,
	}
}

func setup(c *cli.Context) error {
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
	fmt.Fprintf(os.Stdout, "\nInitiate go.mod\n")
	if err := execkit.Run(ctx, &execkit.Command{
		Name:   "go",
		Args:   []string{"mod", "init", pkg},
		Stderr: &stderr,
	}); err != nil {
		return errors.New(stderr.String())
	}
	return nil
}

func createWrapper(p *Param) error {
	fmt.Fprintf(Stdout, "\nCreate wrapper '%s'\n", typicalw)
	return common.ExecuteTmplToFile(typicalw, typicalwTmpl, p)
}

func newProject(p *Param) error {
	projectName := filepath.Base(p.ProjectPkg)

	mkdirAll(p.TypicalTmp)
	mkdirAll("cmd/" + projectName)
	mkdirAll("internal/app")
	mkdirAll("app")

	return nil
}

func mkdirAll(path string) {
	fmt.Fprintf(Stdout, "Mkdir %s\n", path)
	os.MkdirAll(path, 0777)
}
