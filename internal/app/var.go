package app

const (
	typicalw = "typicalw"

	typicalTmpParam    = "typical-tmp"
	projPkgParam       = "project-pkg"
	srcParam           = "src"
	createWrapperParam = "create:wrapper"
)

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
