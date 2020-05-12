package typical

import (
	"github.com/typical-go/typical-go/examples/provide-constructor/helloworld"
	"github.com/typical-go/typical-go/pkg/typgo"
)

// Descriptor of sample
var Descriptor = typgo.Descriptor{
	Name:    "provide-constructor",
	Version: "1.0.0",

	EntryPoint: helloworld.Main2,

	BuildSequences: []interface{}{
		&typgo.StdBuild{},
	},

	Layouts: []string{
		"helloworld",
	},
}
