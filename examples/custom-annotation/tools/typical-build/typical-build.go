package main

import (
	"fmt"
	"time"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "custom-annotation",
	ProjectVersion: "1.0.0",

	Tasks: []typgo.Tasker{
		&typast.AnnotateProject{
			Sources: []string{"internal"},
			Annotators: []typast.Annotator{
				typast.NewAnnotator(printAllAnnotation),
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
		&typgo.RunBinary{Before: typgo.TaskNames{"annotate", "build"}},
	},
}

func printAllAnnotation(c *typast.Context) error {
	fmt.Println("Print all annotation: ")
	for _, a := range c.Annots {
		fmt.Printf("TagName=%s\tName=%s\tType=%T\tParam=%s\tField1=%s\n",
			a.TagName, a.GetName(), a.Decl.Type, a.TagParam, a.TagParam.Get("field1"))
	}
	fmt.Println()
	return nil
}

func main() {
	typgo.Start(&descriptor)
}
