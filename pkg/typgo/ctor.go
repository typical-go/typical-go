package typgo

import (
	"fmt"
	"log"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	ctorTag = "ctor"
)

type (
	// CtorAnnotation represent @ctor annotation
	CtorAnnotation struct{}
)

var _ Action = (*CtorAnnotation)(nil)

// Execute ctor annotation
func (*CtorAnnotation) Execute(c *Context) error {
	var ctors []*typtmpl.Ctor
	for _, annot := range c.ASTStore.Annots {
		if annot.Check(ctorTag, typast.FuncType) {
			ctor, err := typtmpl.CreateCtor(annot)
			if err != nil {
				log.Printf("WARN %s", err.Error())
				continue
			}
			ctors = append(ctors, ctor)
		}
	}

	return writeGoSource(
		fmt.Sprintf("cmd/%s/ctor_annotated.go", c.Descriptor.Name),
		&typtmpl.CtorAnnotated{
			Package: "main",
			Imports: c.Imports,
			Ctors:   ctors,
		},
	)
}
