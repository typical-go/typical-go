package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

// DependencyInjection prebuilding
type DependencyInjection struct{}

var _ Compiler = (*DependencyInjection)(nil)

// Compile dependency-injection
func (d *DependencyInjection) Compile(c *Context) (err error) {
	if err = d.ctor(c); err != nil {
		return
	}

	if err = d.dtor(c); err != nil {
		return
	}

	return
}

func (*DependencyInjection) ctor(c *Context) error {
	var ctors []*typtmpl.Ctor
	for _, annot := range c.ASTStore.Annots {
		ctor, err := CreateCtor(annot)
		if err != nil {
			c.Warnf("ctor: %s", err.Error())
			continue
		}
		if ctor != nil {
			ctors = append(ctors, &typtmpl.Ctor{
				Name: ctor.Param.Name,
				Def:  fmt.Sprintf("%s.%s", ctor.Decl.Pkg, ctor.Decl.Name),
			})
		}
	}

	return writeGoSource(&typtmpl.CtorGenerated{
		Package: "main",
		Imports: c.Imports,
		Ctors:   ctors,
	}, fmt.Sprintf("%s/%s/ctor_generated.go", CmdFolder, c.Descriptor.Name))
}

func (*DependencyInjection) dtor(c *Context) error {
	var dtors []*typtmpl.Dtor
	for _, annot := range c.ASTStore.Annots {
		dtor := CreateDtor(annot)
		if dtor != nil {
			dtors = append(dtors, &typtmpl.Dtor{
				Def: fmt.Sprintf("%s.%s", dtor.Decl.Pkg, dtor.Decl.Name),
			})
		}
	}
	return writeGoSource(&typtmpl.DtorGenerated{
		Package: "main",
		Imports: c.Imports,
		Dtors:   dtors,
	}, fmt.Sprintf("%s/%s/dtor_generated.go", CmdFolder, c.Descriptor.Name))
}
