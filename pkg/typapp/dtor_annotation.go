package typapp

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// DtorAnnotation handle @dtor annotation. No Attributes required.
	DtorAnnotation struct {
		Target string
	}
	// DtorAnnotated template
	DtorAnnotated struct {
		Package string
		Imports []string
		Dtors   []*Dtor
	}
	// Dtor is destructor model
	Dtor struct {
		Def string
	}
)

var _ typannot.Annotator = (*DtorAnnotation)(nil)

// Annotate @dtor
func (a *DtorAnnotation) Annotate(c *typannot.Context) error {
	var dtors []*Dtor
	for _, annot := range c.ASTStore.Annots {
		if annot.CheckFunc(dtorTag) {
			dtors = append(dtors, &Dtor{
				Def: fmt.Sprintf("%s.%s", annot.Decl.Package, annot.Decl.Name),
			})
		}
	}

	target := a.GetTarget(c)
	if err := common.ExecuteTmplToFile(target, dtorAnnotTmpl, &DtorAnnotated{
		Package: "main",
		Imports: c.CreateImports(typgo.ProjectPkg,
			"github.com/typical-go/typical-go/pkg/typapp",
		),
		Dtors: dtors,
	}); err != nil {
		return err
	}
	goImports(target)

	return nil
}

// GetTarget to get generation target for dtor
func (a *DtorAnnotation) GetTarget(c *typannot.Context) string {
	if a.Target == "" {
		a.Target = fmt.Sprintf("cmd/%s/dtor_annotated.go", c.BuildSys.Name)
	}
	return a.Target
}
