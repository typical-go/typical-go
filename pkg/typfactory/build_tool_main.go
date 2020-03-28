package typfactory

import (
	"io"
	"text/template"
)

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

// Write the tyicalw
func (t *BuildToolMain) Write(w io.Writer) (err error) {
	var tmpl *template.Template
	if tmpl, err = template.New("").Parse(buildtoolMain); err != nil {
		return
	}
	return tmpl.Execute(w, t)
}
