package main

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typapp-sample",
	ProjectVersion: "1.0.0",

	Tasks: []typgo.Tasker{
		// generate
		&typgen.Generator{
			Processor: typgen.Processors{
				&typapp.CtorAnnot{},
			},
		},
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{
			Before: typgo.TaskNames{"generate", "build"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
