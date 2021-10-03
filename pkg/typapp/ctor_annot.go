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
	DefaultCtorTag = "@ctor"

	_ typgen.Annotation = (*CtorAnnot)(nil)
)

func (a *CtorAnnot) TagName() string {
	return DefaultCtorTag
}

func (a *CtorAnnot) IsAllowed(d *typgen.Directive) bool {
	return typgen.IsPublic(d)
}

func (a *CtorAnnot) Process(c *typgen.Context) error {
	if len(c.Dirs) < 1 {
		return nil
	}

	c.PutInitSprintf("// <<< [Annotation:%s] ", a.TagName())
	for _, d := range c.Dirs {
		nameTag := d.TagParam.Get("name")
		packagePath := d.PackagePath()

		switch d.Type.(type) {
		case *typgen.Function:
			funcDecl := d.Type.(*typgen.Function)
			if !funcDecl.IsMethod() {
				c.ProvideConstructor(nameTag, packagePath, d.GetName())
			} else {
				a.notSupported(c, d)
			}
		case *typgen.Struct:
			c.PutInit("// TODO: create constructor for struct")
		default:
			a.notSupported(c, d)
		}
	}

	c.PutInitSprintf("// [Annotation:%s] >>>", a.TagName())
	c.PutInit("") // NOTE: intentionally put blank

	return nil
}

func (a *CtorAnnot) notSupported(c *typgen.Context, d *typgen.Directive) {
	c.PutInitSprintf("// '%s' is not supported", d.GetName())
}
