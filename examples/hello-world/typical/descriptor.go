package typical

import (
	"github.com/typical-go/typical-go/examples/hello-world/helloworld"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Modules
var (
	helloApp = helloworld.New()
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "hello-world",
	Version: "0.0.1",

	App: helloApp,

	BuildTool: typbuildtool.New(),
}
