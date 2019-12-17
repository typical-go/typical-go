package tmpl

// ContextWithAppModule template
const ContextWithAppModule = `package typical

import (
	"{{.Pkg}}/app"

	"github.com/typical-go/typical-go/pkg/typcore"
)

// Context of Project
var Context = &typcore.Context{
	Name:      "{{.Name}}",
	Version:   "0.0.1",
	Package:   "{{.Pkg}}",
	AppModule: &app.Module{},
}
`
