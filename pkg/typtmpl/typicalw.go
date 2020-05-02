package typtmpl

import (
	"io"
	"text/template"
)

var _ Template = (*Typicalw)(nil)

const typicalw = `#!/bin/bash

set -e

TYPSRC={{.TypicalSource}}
TYPTMP={{.TypicalTmp}}
TYPGO=$TYPTMP/bin/typical-go

if ! [ -s $TYPGO ]; then
	go build -o $TYPGO $TYPSRC
fi

$TYPGO wrap \
	-typical-tmp=$TYPTMP \
{{if .ProjectPkg }}	-project-pkg="{{.ProjectPkg}}" \
{{end}}
$TYPTMP/bin/build-tool $@
`

// Typicalw writer
type Typicalw struct {
	TypicalTmp    string
	TypicalSource string
	ProjectPkg    string
}

// Execute typicalw template
func (t *Typicalw) Execute(w io.Writer) (err error) {
	var tmpl *template.Template
	if tmpl, err = template.New("").Parse(typicalw); err != nil {
		return
	}
	return tmpl.Execute(w, t)
}
