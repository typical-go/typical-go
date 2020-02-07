package tmpl

// Context template
const Context = `package typical

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typbuild"
)

// Descriptor of {{.Name}}
var Descriptor = typcore.Descriptor{
	Name:      "{{.Name}}",
	Version:   "0.0.1",
	Package:   "{{.Pkg}}",

	Build: typbuild.New(),
}
`
