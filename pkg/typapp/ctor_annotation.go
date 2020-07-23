package typapp

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	ctorTag = "@ctor"
)

type (
	// CtorAnnotation handle @ctor annotation
	// e.g. `@ctor (name:"NAME")`
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

	target := a.GetTarget(c)
	return WriteGoSource(target, &typtmpl.CtorAnnotated{
		Package: "main",
		Imports: c.CreateImports(typgo.ProjectPkg,
			"github.com/typical-go/typical-go/pkg/typapp",
		),
		Ctors: ctors,
	})
}

// GetTarget to return target generation of ctor
func (a *CtorAnnotation) GetTarget(c *typannot.Context) string {
	if a.Target == "" {
		a.Target = fmt.Sprintf("cmd/%s/ctor_annotated.go", c.BuildSys.Name)
	}
	return a.Target
}
