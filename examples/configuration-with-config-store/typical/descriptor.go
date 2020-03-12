package typical

import (
	"github.com/typical-go/typical-go/examples/configuration-with-config-store/server"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Modules
var (
	serverApp = server.New()
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "configuration-with-config-store",
	Version: "0.0.1",

	App: serverApp, // wrap serverApp with Typical App

	BuildTool: typbuildtool.New(),

	Configuration: typcore.NewConfiguration().
		AppendConfigurer(
			serverApp, // Append configurer for the this project
		),
}
