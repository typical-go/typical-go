package main

import (
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "use-config",
		Version: "1.0.0",
		Layouts: []string{"internal"},

		Cmds: []typgo.Cmd{
			&typgo.CompileCmd{
				Action: typgo.Actions{
					&typannot.Annotators{
						&typapp.CfgAnnotation{DotEnv: true},
					},
					&typgo.StdCompile{},
				},
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
