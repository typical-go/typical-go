package typapp

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typgo"
)

type (
	// DtorAnnotation handle @dtor annotation. No Attributes required.
	DtorAnnotation struct {
		Target string
	}
	// DtorTmplData template
	DtorTmplData struct {
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
	dtors := a.CreateDtors(c)
	target := a.GetTarget(c)
	if len(dtors) < 1 {
		os.Remove(target)
		return nil
	}
	data := &DtorTmplData{
		Package: "main",
		Imports: c.CreateImports(typgo.ProjectPkg,
			"github.com/typical-go/typical-go/pkg/typapp",
		),
		Dtors: dtors,
	}
	fmt.Fprintf(Stdout, "Generate @dtor to %s\n", target)
	if err := common.ExecuteTmplToFile(target, dtorAnnotTmpl, data); err != nil {
		return err
	}
	goImports(target)

	return nil
}

// CreateDtors get dtors
func (a *DtorAnnotation) CreateDtors(c *typannot.Context) []*Dtor {
	var dtors []*Dtor
	for _, annot := range c.ASTStore.Annots {
		if annot.CheckFunc(dtorTag) {
			dtors = append(dtors, &Dtor{
				Def: fmt.Sprintf("%s.%s", annot.Decl.Package, annot.Decl.Name),
			})
		}
	}
	return dtors
}

// GetTarget to get generation target for dtor
func (a *DtorAnnotation) GetTarget(c *typannot.Context) string {
	if a.Target == "" {
		a.Target = fmt.Sprintf("cmd/%s/dtor_annotated.go", c.BuildSys.Name)
	}
	return a.Target
}
