package typapp

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	ctorTag = "@ctor"
)

type (
	// CtorAnnotation represent @ctor annotation
	CtorAnnotation struct {
		Target string
	}
)

var _ typannot.Annotator = (*CtorAnnotation)(nil)

// Annotate ctor
func (a *CtorAnnotation) Annotate(c *typannot.Context) error {
	var ctors []*typtmpl.Ctor
	for _, annot := range c.ASTStore.Annots {
		if annot.CheckFunc(ctorTag) {
			ctors = append(ctors, &typtmpl.Ctor{
				Name: annot.TagAttrs.Get("name"),
				Def:  fmt.Sprintf("%s.%s", annot.Decl.Package, annot.Decl.Name),
			})
		}
	}

	return WriteGoSource(
		a.GetTarget(c),
		&typtmpl.CtorAnnotated{Package: "main", Imports: c.Imports, Ctors: ctors},
	)
}

// GetTarget to return target generation of ctor
func (a *CtorAnnotation) GetTarget(c *typannot.Context) string {
	if a.Target == "" {
		a.Target = fmt.Sprintf("cmd/%s/ctor_annotated.go", c.BuildSys.Name)
	}
	return a.Target
}
