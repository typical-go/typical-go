package tmpl

// ContextWithAppModule template
const ContextWithAppModule = `package typical

import (
	"{{.Pkg}}/app"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typbuild"
)

// Descriptor of {{.Name}}
var Descriptor = typcore.Descriptor{
	Name:      "{{.Name}}",
	Version:   "0.0.1",
	Package:   "{{.Pkg}}",

	App: typapp.New(application),

	Build: typbuild.New(),
	
	Configuration: typcfg.New().
		WithConfigure(
			application,
		), 
}

var (
	application = app.New()
)
`
