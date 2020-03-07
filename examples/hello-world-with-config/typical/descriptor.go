package typical

import (
	"github.com/typical-go/typical-go/examples/hello-world-with-config/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Modules
var (
	hello = helloworld.New()
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "hello-world-with-config",
	Version: "0.0.1",

	App: typapp.New(hello),

	BuildTool: typbuildtool.New(),

	Configuration: typcfg.New().
		AppendConfigurer(
			hello, // Append configurer for the this project
		),
}
