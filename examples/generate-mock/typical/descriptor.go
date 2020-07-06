package typical

import (
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "generate-mock",
	Version: "1.0.0",

	Layouts: []string{"internal"},

	Prebuild: &typgo.DependencyInjection{},
	Compile:  &typgo.StdCompile{},
	Run:      &typgo.StdRun{},
	Test:     &typgo.StdTest{},
	Clean:    &typgo.StdClean{},

	Utility: &typmock.Utility{},
}
