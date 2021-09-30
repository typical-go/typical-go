package main

import (
	"time"
	
	"github.com/typical-go/typical-go/pkg/typgen"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typmock"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "my-project",
	ProjectVersion: "0.0.1",

	Tasks: []typgo.Tasker{
		// generate
		&typgen.Generator{
			Annotations: []typgen.Annotation{
				&typapp.CtorAnnot{},
			},
		},
		// build
		&typgo.GoBuild{},
		// test
		&typgo.GoTest{
			Timeout:  30 * time.Second,
			Includes: []string{"internal/*"},
			Excludes: []string{"internal/generated"},
		},
		// run
		&typgo.RunBinary{Before: typgo.TaskNames{"generate", "build"}},
		// mock
		&typmock.GoMock{},
	},
}

func main() {
	typgo.Start(&descriptor)
}
