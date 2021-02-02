package main

import (
	"github.com/typical-go/typical-go/examples/custom-annotation/internal/app"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "custom-annotation",
	ProjectVersion: "1.0.0",

	Tasks: []typgo.Tasker{
		&typast.AnnotateMe{
			Sources: []string{"internal"},
			Annotators: []typast.Annotator{
				&app.MyAnnotation{},
			},
		},
		// test
		&typgo.GoTest{
			Args:     []string{"-timeout=30s"},
			Includes: []string{"internal/*"},
		},
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{Before: typgo.TaskNames{"annotate", "build"}},
	},
}

func main() {
	typgo.Start(&descriptor)
}
