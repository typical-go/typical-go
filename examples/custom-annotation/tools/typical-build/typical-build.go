package main

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "custom-annotation",
	ProjectVersion: "1.0.0",

	Tasks: []typgo.Tasker{
		&typast.AnnotateMe{
			Sources: []string{"internal"},
			Annotators: []typast.Annotator{
				typast.NewAnnotator(func(c *typast.Context) error {
					for _, a := range c.Annots {
						fmt.Printf("TagName=%s\tName=%s\tType=%T\tParam=%s\tField1=%s\n",
							a.TagName, a.GetName(), a.Decl.Type, a.TagParam, a.TagParam.Get("field1"))
					}
					return nil
				}),
			},
		},
		// test
		&typgo.GoTest{
			Args:     []string{"-timeout=30s"},
			Includes: []string{"internal/*"},
		},
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunBinary{Before: typgo.TaskNames{"annotate", "build"}},
	},
}

func main() {
	typgo.Start(&descriptor)
}
