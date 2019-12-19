package tmpl

// ContextWithAppModule template
const ContextWithAppModule = `package typical

import (
	"{{.Pkg}}/app"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of {{.Name}}
var Descriptor = &typcore.ProjectDescriptor{
	Name:      "{{.Name}}",
	Version:   "0.0.1",
	Package:   "{{.Pkg}}",
	AppModule: &app.Module{},
}
`
