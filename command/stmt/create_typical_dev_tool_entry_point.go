package stmt

import (
	"html/template"
	"os"

	"github.com/typical-go/typical-go/typibuild"
)

const typicalDevToolEntryPointTemplate = `package main

import (
	"log"
	"os"

	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typimain"
	"{{ .PackageName }}/typical"
)

func main() {
	cli := typimain.NewTypicalTaskTool(typical.Context)
	err := cli.Cli().Run(os.Args)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
`

type CreateTypicalDevToolEntryPoint struct {
	Project *typibuild.Project
	Target  string
}

func (c CreateTypicalDevToolEntryPoint) Run() (err error) {
	f, err := os.Create(c.Target)
	if err != nil {
		return
	}

	tmpl, _ := template.New("typical_context").Parse(typicalDevToolEntryPointTemplate)
	return tmpl.Execute(f, c.Project)
}
