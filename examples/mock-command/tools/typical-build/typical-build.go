package main

import (
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "mock-command",
		Version: "1.0.0",
		Layouts: []string{"internal"},

		Cmds: []typgo.Cmd{

			&typgo.CompileCmd{
				Before: typannot.Annotators{
					&typapp.CtorAnnotation{},
				},
				Action: &typgo.StdCompile{},
			},

			&typgo.RunCmd{
				Before: typgo.BuildSysRuns{"compile"},
				Action: &typgo.StdRun{},
			},

			&typgo.TestCmd{
				Action: &typgo.StdTest{},
			},

			&typgo.CleanCmd{
				Action: &typgo.StdClean{},
			},

			&typmock.MockCmd{},
		},
	}
)

func main() {
	typgo.Start(&descriptor)
}
