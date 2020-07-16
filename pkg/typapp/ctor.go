package typapp

import (
	"fmt"
	"log"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	ctorTag = "ctor"
)

type (
	// CtorAnnotation represent @ctor annotation
	CtorAnnotation struct {
		Target string
	}
)

var _ typgo.Action = (*CtorAnnotation)(nil)

// Execute ctor annotation
func (a *CtorAnnotation) Execute(c *typgo.Context) error {
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

	return typgo.WriteGoSource(
		a.GetTarget(c),
		&typtmpl.CtorAnnotated{
			Package: "main",
			Imports: c.Imports,
			Ctors:   ctors,
		},
	)
}

// GetTarget to return target generation of ctor
func (a *CtorAnnotation) GetTarget(c *typgo.Context) string {
	if a.Target == "" {
		a.Target = fmt.Sprintf("cmd/%s/ctor_annotated.go", c.Descriptor.Name)
	}
	return a.Target
}
