package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	dtorTag = "dtor"
)

type (
	// DtorAnnotation represent @dtor annotation
	DtorAnnotation struct{}
)

var _ Action = (*DtorAnnotation)(nil)

// Execute @dtor
func (*DtorAnnotation) Execute(c *Context) error {
	var dtors []*typtmpl.Dtor
	for _, annot := range c.ASTStore.Annots {
		if annot.Check(dtorTag, typast.FuncType) {
			dtors = append(dtors, typtmpl.CreateDtor(annot))
		}
	}
	return writeGoSource(
		fmt.Sprintf("cmd/%s/dtor_annotated.go", c.Descriptor.Name),
		&typtmpl.DtorAnnotated{
			Package: "main",
			Imports: c.Imports,
			Dtors:   dtors,
		},
	)
}
