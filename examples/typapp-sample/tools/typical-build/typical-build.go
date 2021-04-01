package main

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typapp-sample",
	ProjectVersion: "1.0.0",

	Tasks: []typgo.Tasker{
		// annotate
		&typast.AnnotateProject{
			Sources: []string{"internal"},
			Annotators: []typast.Annotator{
				&typapp.CtorAnnot{},
			},
		},
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{
			Before: typgo.TaskNames{"annotate", "build"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
