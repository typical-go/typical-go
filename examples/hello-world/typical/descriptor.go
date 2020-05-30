package typical

import (
	"github.com/typical-go/typical-go/examples/hello-world/internal/helloworld"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "hello-world",
	Version: "1.0.0",

	EntryPoint: helloworld.Main,

	Compile: &typgo.StdCompile{},
	Run:     &typgo.StdRun{},
}
