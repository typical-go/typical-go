package app

import (
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

// TypicalwTmpl typicalw template text
var TypicalwTmpl = `#!/bin/bash

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

func wrapper(c *cli.Context) error {
	typicalTmp := getTypicalTmp(c)
	src := getSrc(c)

	projectPkg, err := getProjectPkg(c)
	if err != nil {
		return err
	}

	return common.ExecuteTmplToFile(typicalw, TypicalwTmpl, &Typicalw{
		Src:        src,
		TypicalTmp: typicalTmp,
		ProjectPkg: projectPkg,
	})
}
