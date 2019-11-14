package stmt

import (
	"html/template"
	"os"
)

const appEntryPointTemplate = `package main

import (
	"log"
	"os"

	"{{ .PackageName }}/typical"
	"github.com/typical-go/typical-rest-server/EXPERIMENTAL/typimain"
)
	

func main(){
	app := typimain.NewTypicalApplication(typical.Context)
	err := app.Cli().Run(os.Args)
	if err != nil {
		log.Fatal(err)	
	}
}`

type CreateAppEntryPoint struct {
	PackageName string
	Target      string
}

func (c CreateAppEntryPoint) Run() (err error) {
	f, err := os.Create(c.Target)
	if err != nil {
		return
	}

	tmpl, _ := template.New("typical_context").Parse(appEntryPointTemplate)
	return tmpl.Execute(f, c)
}
