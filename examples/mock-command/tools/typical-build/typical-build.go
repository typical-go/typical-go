package main

import (
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var descriptor = typgo.Descriptor{
	Name:    "mock-command",
	Version: "1.0.0",
	Layouts: []string{"internal"},

	Cmds: []typgo.Cmd{
		// annotate
		&typannot.AnnotateCmd{
			Annotators: []typannot.Annotator{
				&typapp.CtorAnnotation{},
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

		// test
		&typgo.TestCmd{
			Action: &typgo.StdTest{},
		},

		// clean
		&typgo.CleanCmd{
			Action: &typgo.StdClean{},
		},

		// mock
		&typmock.MockCmd{},
	},
}

func main() {
	typgo.Start(&descriptor)
}
