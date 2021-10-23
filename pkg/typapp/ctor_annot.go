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

	_ typgen.Annotator              = (*CtorAnnot)(nil)
	_ typgen.AnnotatedFileProcessor = (*CtorAnnot)(nil)
)

func (a *CtorAnnot) AnnotationName() string {
	return DefaultCtorAnnot
}

func (a *CtorAnnot) IsAllowed(d *typgen.Annotation) bool {
	return typgen.IsPublic(d)
}

func (a *CtorAnnot) ProcessAnnotatedFile(c *typgen.Context, f *typgen.File, annots []*typgen.Annotation) error {

	for _, annot := range annots {
		nameParam := annot.Params.Get("name")
		packagePath := annot.PackagePath()

		switch annot.Decl.Type.(type) {
		case *typgen.Function:
			funcDecl := annot.Decl.Type.(*typgen.Function)
			if !funcDecl.IsMethod() {
				c.ProvideConstructor(nameParam, packagePath, annot.Decl.GetName())
			} else {
				notSupported(c, annot)
			}
		case *typgen.Struct:
			c.AppendInit("// TODO: provide struct constructor")
			c.AppendFileCoder(f, typgen.CodeLine("// TODO: create constructor function for struct\n"))
		case *typgen.Interface:
			c.AppendInit("// TODO: provide interface constructor")
			c.AppendFileCoder(f, typgen.CodeLine("// TODO: create constructor function for interface\n"))
		default:
			notSupported(c, annot)
		}

	}

	return nil
}

func notSupported(c *typgen.Context, annot *typgen.Annotation) {
	c.AppendInitf("// '%s' is not supported", annot.Decl.GetName())
}
