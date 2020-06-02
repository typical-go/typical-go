package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typannot"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

type (
	// DependencyInjection prebuilding
	DependencyInjection struct {
		Configs []*Configuration
	}
)

var _ Prebuilder = (*DependencyInjection)(nil)

// Prebuild dependency-injection
func (d *DependencyInjection) Prebuild(c *Context) (err error) {
	if err = d.ctor(c); err != nil {
		return
	}

	if err = d.dtor(c); err != nil {
		return
	}

	return
}

func (*DependencyInjection) ctor(c *Context) error {
	ctorAnnots, errs := typannot.GetCtors(c.ASTStore)
	for _, a := range ctorAnnots {
		c.Precond.Ctors = append(c.Precond.Ctors, &typtmpl.Ctor{
			Name: a.Param.Name,
			Def:  fmt.Sprintf("%s.%s", a.Decl.Pkg, a.Decl.Name),
		})
	}

	for _, err := range errs {
		c.Warnf("App-Precond: %s", err.Error())
	}
	return nil
}

func (*DependencyInjection) dtor(c *Context) error {
	dtorAnnots, errs := typannot.GetDtors(c.ASTStore)
	for _, a := range dtorAnnots {
		c.Precond.Dtors = append(c.Precond.Dtors, &typtmpl.Dtor{
			Def: fmt.Sprintf("%s.%s", a.Decl.Pkg, a.Decl.Name),
		})
	}

	for _, err := range errs {
		c.Warnf("App-Precond: %s", err.Error())
	}
	return nil
}
