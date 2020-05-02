package typtmpl

import (
	"io"
	"text/template"
)

var _ Template = (*BuildToolMain)(nil)

const buildtoolMain = `package main

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"{{.DescPkg}}"
)

func main() {
	typcore.LaunchBuildTool(&typical.Descriptor)
}
`

// BuildToolMain is writer to generate main.go for app
type BuildToolMain struct {
	DescPkg string
}

// Execute build tool main template
func (t *BuildToolMain) Execute(w io.Writer) (err error) {
	var tmpl *template.Template
	if tmpl, err = template.New("").Parse(buildtoolMain); err != nil {
		return
	}
	return tmpl.Execute(w, t)
}
