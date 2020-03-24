package typical

import (
	"github.com/typical-go/typical-go/examples/configuration-with-invocation/server"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Modules
var (
	serverApp = server.New()
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "configuration-with-invocation",
	Version: "1.0.0",

	App: typapp.
		Create(serverApp), // wrap serverApp with Typical App

	BuildTool: typbuildtool.
		Create(
			typbuildtool.StandardBuild(),
		),

	ConfigManager: typcfg.
		Create(
			serverApp, // Append configurer for the this project
		),
}
