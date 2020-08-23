package main

import (
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

var mainPkg = "."

var descriptor = typgo.Descriptor{
	ProjectName:    "typical-go",
	ProjectVersion: "0.10.8",
	ProjectLayouts: []string{"internal", "pkg"},

	Cmds: []typgo.Cmd{
		// compile
		&typgo.CompileProject{MainPackage: mainPkg},
		// test
		&typgo.TestProject{},
		// clean
		&typgo.CleanProject{},
		// run
		&typgo.RunCmd{
			Before: typgo.BuildCmdRuns{"compile"},
			Action: &typgo.RunProject{},
		},
		// examples
		&typgo.Command{
			Name:    "examples",
			Aliases: []string{"e"},
			Usage:   "Test all example",
			Action:  typgo.NewAction(testExamples),
		},
		// release
		&typrls.ReleaseCmd{
			Before: typgo.BuildCmdRuns{"test", "examples"},
			Action: &typrls.ReleaseProject{
				Validator: typrls.DefaultValidator,
				Releaser:  &typrls.Github{Owner: "typical-go", Repo: "typical-go"},
			},
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
