package tmpl

// ContextWithAppModule template
const ContextWithAppModule = `package typical

import (
	"{{.Pkg}}/app"

	"github.com/typical-go/typical-go/pkg/typcore"
)

var (
	application = app.New()
	
	// Descriptor of {{.Name}}
	Descriptor = typcore.ProjectDescriptor{
		Name:      "{{.Name}}",
		Version:   "0.0.1",
		Package:   "{{.Pkg}}",

		App: typcore.NewApp(application),
		
		Configuration: typcore.NewConfiguration().
			WithConfigure(
				application,
			), 
	}
)
`
