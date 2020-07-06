package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

// Descriptor of typical-go
var Descriptor = typgo.Descriptor{
	Name:    "typical-go",
	Version: "0.9.57",

	Layouts: []string{"wrapper", "pkg"},

	Test:    &typgo.StdTest{},
	Compile: &typgo.StdCompile{},
	Run:     &typgo.StdRun{},
	Clean:   &typgo.StdClean{},
	Release: &typgo.Github{Owner: "typical-go", RepoName: "typical-go"},

	Utility: typgo.CreateUtility(&cli.Command{
		Name:    "examples",
		Aliases: []string{"e"},
		Usage:   "Test all example",
		Action: func(cliCtx *cli.Context) (err error) {
			gotest := &execkit.GoTest{
				Targets: []string{"./examples/..."},
			}

			cmd := gotest.Command()
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			return cmd.Run(cliCtx.Context)
		},
	}),
}
