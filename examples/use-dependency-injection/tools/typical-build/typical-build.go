package main

import (
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "use-dependency-injection",
	ProjectVersion: "1.0.0",
	ProjectLayouts: []string{"internal"},

	Cmds: []typgo.Cmd{
		// annotate
		&typannot.AnnotateCmd{
			Annotators: []typannot.Annotator{
				&typapp.CtorAnnotation{},
				&typapp.DtorAnnotation{},
			},
		},

		// compile
		&typgo.CompileCmd{
			Action: &typgo.StdCompile{},
		},

		// run
		&typgo.RunCmd{
			Before: typgo.BuildSysRuns{"annotate", "compile"},
			Action: &typgo.StdRun{},
		},

		// clean
		&typgo.CleanCmd{
			Action: &typgo.StdClean{},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
