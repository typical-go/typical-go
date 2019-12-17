package tmpl

// Context template
const Context = `package typical

import (
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Context of Project
var Context = &typcore.Context{
	Name:      "{{.Name}}",
	Version:   "0.0.1",
	Package:   "{{.Pkg}}",
}
`
