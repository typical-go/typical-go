package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typtmpl"
)

type (
	// DependencyInjection prebuilding
	DependencyInjection struct{}
)

var _ Prebuilder = (*DependencyInjection)(nil)

// Prebuild dependency-injection
func (d *DependencyInjection) Prebuild(c *PrebuildContext) (err error) {
	if err = d.ctor(c); err != nil {
		return
	}

	if err = d.dtor(c); err != nil {
		return
	}

	return
}

func (*DependencyInjection) ctor(c *PrebuildContext) error {
	for _, annot := range c.ASTStore.Annots {
		ctor, err := CreateCtor(annot)
		if err != nil {
			c.Warnf("ctor: %s", err.Error())
			continue
		}
		if ctor != nil {
			c.Precond.Ctors = append(c.Precond.Ctors, &typtmpl.Ctor{
				Name: ctor.Param.Name,
				Def:  fmt.Sprintf("%s.%s", ctor.Decl.Pkg, ctor.Decl.Name),
			})
		}
	}

	return nil
}

func (*DependencyInjection) dtor(c *PrebuildContext) error {
	for _, annot := range c.ASTStore.Annots {
		dtor := CreateDtor(annot)
		if dtor != nil {
			c.Precond.Dtors = append(c.Precond.Dtors, &typtmpl.Dtor{
				Def: fmt.Sprintf("%s.%s", dtor.Decl.Pkg, dtor.Decl.Name),
			})
		}
	}
	return nil
}
