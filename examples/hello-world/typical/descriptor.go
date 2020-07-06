package typical

import (
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "hello-world",
	Version: "1.0.0",

	Compile: &typgo.StdCompile{},
	Run:     &typgo.StdRun{},
	Clean:   &typgo.StdClean{},
}
