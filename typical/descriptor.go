package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/wrapper"
	"github.com/urfave/cli/v2"
)

// Descriptor of typical-go
var Descriptor = typgo.Descriptor{
	Name:    "typical-go",
	Version: "0.9.56",

	EntryPoint: wrapper.Main,
	Layouts:    []string{"wrapper", "pkg"},

	Test:    &typgo.StdTest{},
	Compile: &typgo.StdCompile{},
	Run:     &typgo.StdRun{},
	Clean:   &typgo.StdClean{},
	Release: &typgo.Github{Owner: "typical-go", RepoName: "typical-go"},

	Utility: typgo.CreateUtility(&cli.Command{
		Name:    "test-example",
		Aliases: []string{"e"},
		Usage:   "Test all example",
		Action: func(cliCtx *cli.Context) (err error) {
			gotest := &buildkit.GoTest{
				Targets: []string{"./examples/..."},
			}

			cmd := gotest.Command()
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			return cmd.Run(cliCtx.Context)
		},
	}), // Test all the examples
}
