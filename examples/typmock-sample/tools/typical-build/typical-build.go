package main

import (
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "typmock-sample",
	ProjectVersion: "1.0.0",

	Tasks: []typgo.Tasker{
		// mock
		&typmock.GenerateMock{
			Sources: []string{"internal"},
		},
		// test
		&typgo.GoTest{
			Args:     []string{"-timeout=30s"},
			Includes: []string{"internal/*"},
			Excludes: []string{"**/*_mock"},
		},
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{Before: typgo.TaskNames{"build"}},
	},
}

func main() {
	typgo.Start(&descriptor)
}
