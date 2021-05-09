package main

import (
	"fmt"
	"time"

	"github.com/typical-go/typical-go/pkg/typgen"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "custom-annotation",
	ProjectVersion: "1.0.0",

	Tasks: []typgo.Tasker{
		&typgen.Generator{
			Processor: typgen.Processors{
				&typgen.Annotation{
					ProcessFn: printAllAnnotation,
				},
			},
		},
		// test
		&typgo.GoTest{
			Timeout:  30 * time.Second,
			Includes: []string{"internal/*"},
		},
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{Before: typgo.TaskNames{"generate", "build"}},
	},
}

func printAllAnnotation(c *typgo.Context, directives typgen.Directives) error {
	fmt.Println("Print all annotation: ")
	for _, a := range directives {
		fmt.Printf("TagName=%s\tName=%s\tType=%T\tParam=%s\tField1=%s\n",
			a.TagName, a.GetName(), a.Decl.Type, a.TagParam, a.TagParam.Get("field1"))
	}
	fmt.Println()
	return nil
}

func main() {
	typgo.Start(&descriptor)
}
