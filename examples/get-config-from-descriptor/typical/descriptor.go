package typical

import (
	"github.com/typical-go/typical-go/examples/get-config-from-descriptor/server"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "get-config-from-descriptor",
	Version: "1.0.0",

	App: typcore.NewApp(server.Main), // wrap serverApp with Typical App

	BuildTool: typbuildtool.
		Create(
			typbuildtool.StandardBuild(),
		),

	ConfigManager: typcfg.
		Create(
			typcfg.NewConfiguration(server.ConfigName, &server.Config{}), // register serverApp configurer
		),
}
