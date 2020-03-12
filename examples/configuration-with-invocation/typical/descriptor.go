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
	Version: "0.0.1",

	App: typapp.New(serverApp), // wrap serverApp with Typical App

	BuildTool: typbuildtool.New(),

	Configuration: typcfg.New().
		AppendConfigurer(
			serverApp, // Append configurer for the this project
		),
}
