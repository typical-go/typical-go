package app

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/urfave/cli/v2"
)

type (
	// Typicalw writer
	Typicalw struct {
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
		},
		Action: wrapper,
	}
}

func wrapper(c *cli.Context) error {
	typicalTmp := getTypicalTmp(c)
	src := getSrc(c)

	projectPkg, err := getProjectPkg(c)
	if err != nil {
		return err
	}

	fmt.Fprintf(Stdout, "Create wrapper '%s'\n", typicalw)
	return common.ExecuteTmplToFile(typicalw, typicalwTmpl, &Typicalw{
		Src:        src,
		TypicalTmp: typicalTmp,
		ProjectPkg: projectPkg,
	})
}
