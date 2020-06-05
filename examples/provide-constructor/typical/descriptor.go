package typical

import (
	"github.com/typical-go/typical-go/examples/provide-constructor/internal/helloworld"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "provide-constructor",
	Version: "1.0.0",

	EntryPoint: helloworld.Main2,
	Layouts:    []string{"internal"},

	Prebuild: &typgo.DependencyInjection{},
	Compile:  &typgo.StdCompile{},
	Run:      &typgo.StdRun{},
	Clean:    &typgo.StdClean{},
}
