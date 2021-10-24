package typapp

import (
	"fmt"

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
				c.ProvideConstructor(nameParam, packagePath, funcDecl.GetName())
			} else {
				notSupported(c, annot)
			}
		case *typgen.Struct:
			stDecl := annot.Decl.Type.(*typgen.Struct)
			name := fmt.Sprintf("New%s", stDecl.GetName())
			c.AppendInit("// TODO: provide struct constructor")
			c.AppendFileCoder(f, &typgen.Function{
				Name: name,
				Body: typgen.CodeLines{
					"// TODO: create constructor function for struct",
				},
			})
		case *typgen.Interface:
			itDecl := annot.Decl.Type.(*typgen.Interface)
			name := fmt.Sprintf("New%s", itDecl.GetName())
			c.AppendInit("// TODO: provide interface constructor")
			c.AppendFileCoder(f, &typgen.Function{
				Name: name,
				Body: typgen.CodeLines{
					"// TODO: create constructor function for interface",
				},
			})
		default:
			notSupported(c, annot)
		}

	}

	return nil
}

func notSupported(c *typgen.Context, annot *typgen.Annotation) {
	c.AppendInitf("// '%s' is not supported", annot.Decl.GetName())
}
