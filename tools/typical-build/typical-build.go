package main

import (
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

var mainPkg = "."

var descriptor = typgo.Descriptor{
	ProjectName:    "typical-go",
	ProjectVersion: "0.10.13",
	ProjectLayouts: []string{"internal", "pkg"},

	Cmds: []typgo.Cmd{
		// compile
		&typgo.CompileProject{MainPackage: mainPkg},
		// test
		&typgo.TestProject{},
		// run
		&typgo.RunProject{
			Before: typgo.BuildCmdRuns{"compile"},
		},
		// examples
		&typgo.Command{
			Name:    "examples",
			Aliases: []string{"e"},
			Usage:   "Test all example",
			Action:  typgo.NewAction(testExamples),
		},
		// release
		&typrls.ReleaseProject{
			Before: typgo.BuildCmdRuns{"test", "examples"},
			// Releaser: &typrls.CrossCompiler{
			// 	Targets:     []typrls.Target{"darwin/amd64", "linux/amd64"},
			// 	MainPackage: mainPkg,
			// },
			Publisher: &typrls.Github{Owner: "typical-go", Repo: "typical-go"},
		},
	},
}

func testExamples(c *typgo.Context) error {
	return c.Execute(&execkit.GoTest{
		Packages: []string{"./examples/..."},
	})
}

func main() {
	typgo.Start(&descriptor)
}
