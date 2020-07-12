package typtmpl

import (
	"io"
)

// Typicalw writer
type Typicalw struct {
	Src        string
	TypicalTmp string
	ProjectPkg string
}

var _ Template = (*Typicalw)(nil)

const typicalw = `#!/bin/bash

set -e

TYPTMP={{.TypicalTmp}}
TYPGO=$TYPTMP/bin/typical-go

if ! [ -s $TYPGO ]; then
	go build -o $TYPGO github.com/typical-go/typical-go/cmd/typical-go
fi

$TYPGO \
	-src="{{.Src}}" \
	-project-pkg="{{.ProjectPkg}}" \
	-typical-tmp=$TYPTMP \
	$@
`

// Execute typicalw template
func (t *Typicalw) Execute(w io.Writer) (err error) {
	return Execute("typicalw", typicalw, t, w)
}
