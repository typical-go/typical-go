package main

import (
	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	AppName:    "typapp-sample",
	AppVersion: "1.0.0",
	AppLayouts: []string{"internal"},

	Cmds: []typgo.Cmd{
		// annotate
		&typast.AnnotateCmd{
			Annotators: []typast.Annotator{
				&typapp.CtorAnnotation{},
				&typapp.DtorAnnotation{},
			},
		},
		// compile
		&typgo.CompileProject{},
		// run
		&typgo.RunCmd{
			Before: typgo.BuildCmdRuns{"annotate", "compile"},
			Action: &typgo.RunProject{},
		},
		// clean
		&typgo.CleanProject{},
	},
}

func main() {
	typgo.Start(&descriptor)
}
