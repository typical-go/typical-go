package main

import (
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typmock-sample",
	ProjectVersion: "1.0.0",
	ProjectLayouts: []string{"internal"},

	Cmds: []typgo.Cmd{
		// compile
		&typgo.CompileProject{},
		// run
		&typgo.RunProject{
			Before: typgo.BuildCmdRuns{"compile"},
		},
		// test
		&typgo.GoTest{
			Args:     []string{"-timeout=30s"},
			Includes: []string{"internal/*"},
			Excludes: []string{"**/*_mock"},
		},
		// mock
		&typmock.MockCmd{},
	},
}

func main() {
	typgo.Start(&descriptor)
}
