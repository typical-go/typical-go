package typapp

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// CtorAnnotation handle @ctor annotation
	// e.g. `@ctor (name:"NAME")`
	CtorAnnotation struct {
		Target string
	}
	// CtorTmplData template
	CtorTmplData struct {
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
	ctors := a.CreateCtors(c)
	target := a.GetTarget(c)
	if len(ctors) < 1 {
		os.Remove(target)
		return nil
	}
	data := &CtorTmplData{
		Package: "main",
		Imports: c.CreateImports(typgo.ProjectPkg,
			"github.com/typical-go/typical-go/pkg/typapp",
		),
		Ctors: ctors,
	}
	fmt.Fprintf(Stdout, "Generate @ctor to %s\n", target)
	if err := common.ExecuteTmplToFile(target, ctorAnnotTmpl, data); err != nil {
		return err
	}
	goImports(target)
	return nil
}

// CreateCtors get ctors
func (a *CtorAnnotation) CreateCtors(c *typannot.Context) []*Ctor {
	var ctors []*Ctor
	for _, annot := range c.ASTStore.Annots {
		if annot.CheckFunc(ctorTag) {
			ctors = append(ctors, &Ctor{
				Name: annot.TagParam.Get("name"),
				Def:  fmt.Sprintf("%s.%s", annot.Decl.Package, annot.Decl.Name),
			})
		}
	}
	return ctors
}

// GetTarget to return target generation of ctor
func (a *CtorAnnotation) GetTarget(c *typannot.Context) string {
	if a.Target == "" {
		a.Target = fmt.Sprintf("cmd/%s/ctor_annotated.go", c.BuildSys.Name)
	}
	return a.Target
}
