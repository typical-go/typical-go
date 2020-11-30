package main

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typapp-sample",
	ProjectVersion: "1.0.0",
	ProjectLayouts: []string{"internal"},

	Tasks: []typgo.Tasker{
		// annotate
		&typast.AnnotateProject{
			Annotators: []typast.Annotator{
				&typapp.CtorAnnotation{},
			},
		},
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{
			Before: typgo.BuildCmdRuns{"annotate", "compile"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
