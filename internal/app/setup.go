package app

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/urfave/cli/v2"
)

type (
	typicalwTmplData struct {
		Src        string
		TypicalTmp string
		ProjectPkg string
	}
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
	-src="{{.Src}}" \
	-project-pkg="{{.ProjectPkg}}" \
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
		},
		Action: wrapper,
	}
}

func wrapper(c *cli.Context) error {
	if gomod := c.String("gomod"); gomod != "" {
		var stderr strings.Builder
		fmt.Fprintf(os.Stdout, "\nInitiate go.mod\n")
		if err := execkit.Run(c.Context, &execkit.Command{
			Name:   "go",
			Args:   []string{"mod", "init", gomod},
			Stderr: &stderr,
		}); err != nil {
			return errors.New(stderr.String())
		}
	}

	typicalTmp := getTypicalTmp(c)
	src := getSrc(c)
	projectPkg, err := getProjectPkg(c)
	if err != nil {
		return err
	}

	fmt.Fprintf(Stdout, "\nCreate wrapper '%s'\n", typicalw)
	return common.ExecuteTmplToFile(typicalw, typicalwTmpl, &typicalwTmplData{
		Src:        src,
		TypicalTmp: typicalTmp,
		ProjectPkg: projectPkg,
	})
}
