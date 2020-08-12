package main

import (
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

var mainPackage = "."

var descriptor = typgo.Descriptor{
	ProjectName:    "typical-go",
	ProjectVersion: "0.10.4",
	ProjectLayouts: []string{"internal", "pkg"},

	Cmds: []typgo.Cmd{

		// compile
		&typgo.CompileCmd{
			Action: &typgo.StdCompile{MainPackage: mainPackage},
		},

		// run
		&typgo.RunCmd{
			Before: typgo.BuildSysRuns{"compile"},
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

		// examples
		&typgo.Command{
			Name:    "examples",
			Aliases: []string{"e"},
			Usage:   "Test all example",
			Action:  typgo.NewAction(testExamples),
		},

		// release
		&typrls.ReleaseCmd{
			Before:     typgo.BuildSysRuns{"test", "examples"},
			Validation: typrls.DefaultValidation,
			Summary:    typrls.DefaultSummary,

			Releaser: typrls.Releasers{
				&typrls.Compile{
					MainPackage: mainPackage,
					Targets:     []typrls.Target{"darwin/amd64", "linux/amd64"},
				},
				&typrls.Github{Owner: "typical-go", Repo: "typical-go"},
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
