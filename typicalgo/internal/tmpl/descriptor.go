package tmpl

// Descriptor template
const Descriptor = `package typical

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
)

// Descriptor of {{.Name}}
var Descriptor = typcore.Descriptor{
	Name:      "{{.Name}}",
	Version:   "0.0.1",
	Package:   "{{.Pkg}}",

	Build: typbuildtool.New(),
}
`
