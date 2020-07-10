package typical

import (
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "provide-constructor",
	Version: "1.0.0",

	Layouts: []string{"internal"},

	Compile: &typgo.StdCompile{
		Before: &typgo.DependencyInjection{},
	},
	Run:   &typgo.StdRun{},
	Clean: &typgo.StdClean{},
}
