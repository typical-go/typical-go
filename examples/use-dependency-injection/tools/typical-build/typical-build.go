package main

import (
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "use-dependency-injection",
		Version: "1.0.0",
		Layouts: []string{"internal"},

		Cmds: []typgo.Cmd{
			&typgo.CompileCmd{
				Before: &typannot.Annotators{
					&typapp.CtorAnnotation{},
					&typapp.DtorAnnotation{},
				},
				Action: &typgo.StdCompile{},
			},
			&typgo.RunCmd{
				Before: typgo.BuildSysRuns{"compile"},
				Action: &typgo.StdRun{},
			},
			&typgo.CleanCmd{
				Action: &typgo.StdClean{},
			},
		},
	}
)

func main() {
	typgo.Start(&descriptor)
}
