package main

import (
	"os"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typrls"
)

var mainPkg = "."

var descriptor = typgo.Descriptor{
	ProjectName:    "typical-go",
	ProjectVersion: "0.10.16",
	ProjectLayouts: []string{"internal", "pkg"},

	Cmds: []typgo.CliCommander{
		// compile
		&typgo.GoBuild{MainPackage: mainPkg},
		// test
		&typgo.GoTest{
			Args:     []string{"-timeout=30s"},
			Includes: []string{"internal/*", "pkg/*"},
		},
		// run
		&typgo.RunProject{
			Before: typgo.BuildCmdRuns{"compile"},
		},
		// examples
		&typgo.Command{
			Name:    "examples",
			Aliases: []string{"e"},
			Usage:   "Test all example",
			Action: typgo.NewAction(func(c *typgo.Context) error {
				return c.Execute(&execkit.Command{
					Name:   "go",
					Args:   []string{"test", "./examples/..."},
					Stdout: os.Stdout,
					Stderr: os.Stderr,
				})
			}),
		},
		// release
		&typrls.ReleaseProject{
			Before: typgo.BuildCmdRuns{"test", "examples"},
			Releaser: &typrls.CrossCompiler{
				Targets:     []typrls.Target{"darwin/amd64", "linux/amd64"},
				MainPackage: mainPkg,
			},
			// Publisher: &typrls.Github{Owner: "typical-go", Repo: "typical-go"},
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
