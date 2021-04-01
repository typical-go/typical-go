package main

import (
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "my-project",
	ProjectVersion: "0.0.1",

	Tasks: []typgo.Tasker{
		// annotate
		&typast.AnnotateProject{
			Sources: []string{"internal"},
			Annotators: []typast.Annotator{
				&typapp.CtorAnnot{},
			},
		},
		// build
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{Before: typgo.TaskNames{"annotate", "build"}},
		// mock
		&typmock.GoMock{
			Sources: []string{"internal"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
