package generator

import (
	"html/template"
	"os"
)

const readmeText = `# {{.ProjectTitle}} 

{{.ProjectDescription}}
`

type ReadMe struct {
	ProjectTitle       string
	ProjectDescription string
	ProjectPath        string
}

func (r ReadMe) Generate(path string) (err error) {
	tmpl, _ := template.New("readme").Parse(readmeText)
	tmpl.Execute(os.Stdout, r)

	return
}
