package main

import (
	"log"

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
				Action: &typgo.Actions{
					&typannot.Annotators{
						&typapp.CtorAnnotation{},
						&typapp.DtorAnnotation{},
					},
					&typgo.StdCompile{},
				},
			},
			&typgo.RunCmd{
				Precmds: []string{"compile"},
				Action:  &typgo.StdRun{},
			},
			&typgo.CleanCmd{
				Action: &typgo.StdClean{},
			},
		},
	}
)

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
