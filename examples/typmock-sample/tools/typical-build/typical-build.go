package main

import (
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typmock-sample",
	ProjectVersion: "1.0.0",
	ProjectLayouts: []string{"internal"},

	Cmds: []typgo.Cmd{
		// annotate
		&typast.AnnotateCmd{
			Annotators: []typast.Annotator{
				&typapp.CtorAnnotation{},
			},
		},
		// compile
		&typgo.CompileProject{},
		// run
		&typgo.RunCmd{
			Before: typgo.BuildCmdRuns{"annotate", "compile"},
			Action: &typgo.RunProject{},
		},
		// test
		&typgo.TestProject{},
		// clean
		&typgo.CleanProject{},
		// mock
		&typmock.MockCmd{},
	},
}

func main() {
	typgo.Start(&descriptor)
}
