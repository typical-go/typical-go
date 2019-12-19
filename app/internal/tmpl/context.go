package tmpl

// Context template
const Context = `package typical

import (
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of {{.Name}}
var Descriptor = &typcore.ProjectDescriptor{
	Name:      "{{.Name}}",
	Version:   "0.0.1",
	Package:   "{{.Pkg}}",
}
`
