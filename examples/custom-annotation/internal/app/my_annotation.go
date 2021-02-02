package app

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typast"
)

// MyAnnotation ...
type (
	MyAnnotation struct{}
)

var _ typast.Annotator = (*MyAnnotation)(nil)

// Annotate the project
func (*MyAnnotation) Annotate(c *typast.Context) error {
	for _, a := range c.Annots {
		fmt.Printf("TagName=%s\tName=%s\tType=%T\tParam=%s\tField1=%s\n",
			a.TagName, a.GetName(), a.Decl.Type, a.TagParam, a.TagParam.Get("field1"))
	}
	return nil
}
