package stmt

import (
	"html/template"
	"os"

	"github.com/typical-go/typical-go/typibuild"
)

const typicalContextTemplate = `package typical

import(
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typictx"
)

// Context instance of Context
var Context typictx.Context

func init() {
	Context = typictx.Context{
		Name:        "{{ .Name }}",
		Version:     "{{ .Version }}",
		Description: "{{ .Description }}",
	}
}
`

// CreateTypicalContext to create Typical Context in Target file
type CreateTypicalContext struct {
	Project *typibuild.Project
	Target  string
}

// Run the create typical context
func (c CreateTypicalContext) Run() (err error) {
	f, err := os.Create(c.Target)
	if err != nil {
		return
	}

	tmpl, _ := template.New("typical_context").Parse(typicalContextTemplate)
	return tmpl.Execute(f, c.Project)
}
