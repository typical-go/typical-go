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

func (a *CtorAnnot) Process(c *typgen.Context) error {
	if len(c.Annotations) < 1 {
		return nil
	}

	c.PutInitSprintf("// <<< [Annotator:%s] ", a.AnnotationName())
	for _, annot := range c.Annotations {
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
			c.PutInit("// TODO: create constructor for struct")
		default:
			a.notSupported(c, annot)
		}
	}

	c.PutInitSprintf("// [Annotator:%s] >>>", a.AnnotationName())
	c.PutInit("") // NOTE: intentionally put blank

	return nil
}

func (a *CtorAnnot) notSupported(c *typgen.Context, annot *typgen.Annotation) {
	c.PutInitSprintf("// '%s' is not supported", annot.Decl.GetName())
}
