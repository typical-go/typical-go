package typapp

import (
	"github.com/typical-go/typical-go/pkg/typgen"
)

type (
	// CtorAnnot handle @ctor annotation
	// e.g. `@ctor (name:"NAME")`
	CtorAnnot struct{}
)

var (
	DefaultCtorAnnot = "@ctor"

	_ typgen.Annotator = (*CtorAnnot)(nil)
)

func (a *CtorAnnot) AnnotationName() string {
	return DefaultCtorAnnot
}

func (a *CtorAnnot) IsAllowed(d *typgen.Annotation) bool {
	return typgen.IsPublic(d)
}

func (a *CtorAnnot) ProcessAnnot(c *typgen.Context, annot *typgen.Annotation) error {
	nameParam := annot.Params.Get("name")
	packagePath := annot.PackagePath()

	switch annot.Decl.Type.(type) {
	case *typgen.Function:
		funcDecl := annot.Decl.Type.(*typgen.Function)
		if !funcDecl.IsMethod() {
			c.ProvideConstructor(nameParam, packagePath, annot.Decl.GetName())
		} else {
			a.notSupported(c, annot)
		}
	case *typgen.Struct:
		c.AppendInit("// TODO: create constructor for struct")
	default:
		a.notSupported(c, annot)
	}

	return nil
}

func (a *CtorAnnot) notSupported(c *typgen.Context, annot *typgen.Annotation) {
	c.AppendInitf("// '%s' is not supported", annot.Decl.GetName())
}
