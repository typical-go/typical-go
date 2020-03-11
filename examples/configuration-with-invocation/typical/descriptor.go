package typical

import (
	"github.com/typical-go/typical-go/examples/configuration-with-invocation/helloworld"
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
	Name:    "configuration-with-invocation",
	Version: "0.0.1",

	App: typapp.New(hello),

	BuildTool: typbuildtool.New(),

	Configuration: typcfg.New().
		AppendConfigurer(
			hello, // Append configurer for the this project
		),
}