package typical

import (
	"github.com/typical-go/typical-go/examples/get-config-from-descriptor/server"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Modules
var (
	serverApp = server.New()
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "get-config-from-descriptor",
	Version: "0.0.1",

	App: serverApp, // wrap serverApp with Typical App

	BuildTool: typbuildtool.New(),

	ConfigManager: typcore.NewConfigManager().
		WithConfigurer(
			serverApp, // register serverApp configurer
		),
}
