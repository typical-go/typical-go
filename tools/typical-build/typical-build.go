package main

import (
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

var mainPkg = "."

var descriptor = typgo.Descriptor{
	AppName:    "typical-go",
	AppVersion: "0.10.11",
	AppLayouts: []string{"internal", "pkg"},

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
		&typrls.ReleaseProject{
			Before:    typgo.BuildCmdRuns{"test", "examples"},
			Validator: typrls.DefaultValidator,
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
