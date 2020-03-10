package tmpl

type ConstructorData struct {
	Imports      []string
	Constructors []string
}

const Constructor = `package typical

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typdep"
	{{range $import := .Imports}}"{{$import}}"
	{{end}}
)

func init() {
	{{if .Constructors}}typapp.AppendConstructor({{range $constructor := .Constructors}}
		typdep.NewConstructor({{$constructor}}),{{end}}
	)
{{end}}}`
