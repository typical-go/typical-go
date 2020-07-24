package typapp

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
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
	// CtorAnnotated template
	CtorAnnotated struct {
		Package string
		Imports []string
		Ctors   []*Ctor
	}
	// Ctor is constructor model
	Ctor struct {
		Name string `json:"name"`
		Def  string `json:"-"`
	}
)

var _ typannot.Annotator = (*CtorAnnotation)(nil)

// Annotate ctor
func (a *CtorAnnotation) Annotate(c *typannot.Context) error {
	var ctors []*Ctor
	for _, annot := range c.ASTStore.Annots {
		if annot.CheckFunc(ctorTag) {
			ctors = append(ctors, &Ctor{
				Name: annot.TagParam.Get("name"),
				Def:  fmt.Sprintf("%s.%s", annot.Decl.Package, annot.Decl.Name),
			})
		}
	}
	target := a.GetTarget(c)
	if err := common.ExecuteTmplToFile(target, ctorAnnotTmpl, &CtorAnnotated{
		Package: "main",
		Imports: c.CreateImports(typgo.ProjectPkg,
			"github.com/typical-go/typical-go/pkg/typapp",
		),
		Ctors: ctors,
	}); err != nil {
		return err
	}
	goImports(target)
	return nil
}

// GetTarget to return target generation of ctor
func (a *CtorAnnotation) GetTarget(c *typannot.Context) string {
	if a.Target == "" {
		a.Target = fmt.Sprintf("cmd/%s/ctor_annotated.go", c.BuildSys.Name)
	}
	return a.Target
}
