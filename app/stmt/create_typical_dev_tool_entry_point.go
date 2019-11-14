package stmt

import (
	"html/template"
	"os"
)

const typicalDevToolEntryPointTemplate = `package main

import (
	"log"
	"os"

	"github.com/typical-go/typical-go/pkg/typimain"
	"{{ .PackageName }}/typical"
)

func main() {
	cli := typimain.NewTypicalDevTool(typical.Context)
	err := cli.Cli().Run(os.Args)
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
`

type CreateTypicalDevToolEntryPoint struct {
	PackageName string
	Target      string
}

func (c CreateTypicalDevToolEntryPoint) Run() (err error) {
	f, err := os.Create(c.Target)
	if err != nil {
		return
	}

	tmpl, _ := template.New("typical_context").Parse(typicalDevToolEntryPointTemplate)
	return tmpl.Execute(f, c)
}
