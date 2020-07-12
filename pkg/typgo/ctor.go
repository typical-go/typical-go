package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typast"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

var (
	ctorTags = []string{"ctor"}
)

type (
	// CtorAnnotation represent @ctor annotation
	CtorAnnotation struct{}
	// Ctor is contructor annotation
	Ctor struct {
		*typast.Annot
		Param CtorParam
	}
	// CtorParam is parameter for ctor annotation
	CtorParam struct {
		Name string `json:"name"`
	}
)

var _ Compiler = (*CtorAnnotation)(nil)

// Compile ctor annotation
func (*CtorAnnotation) Compile(c *Context) error {
	var ctors []*typtmpl.Ctor
	for _, annot := range c.ASTStore.Annots {
		ctor, err := ParseCtor(annot)
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

	return writeGoSource(
		fmt.Sprintf("cmd/%s/ctor_annotated.go", c.Descriptor.Name),
		&typtmpl.CtorAnnotated{
			Package: "main",
			Imports: c.Imports,
			Ctors:   ctors,
		},
	)
}

// ParseCtor annotation
func ParseCtor(annot *typast.Annot) (*Ctor, error) {
	if !IsFuncTag(annot, ctorTags...) {
		return nil, nil
	}

	var ctor Ctor
	if err := annot.Unmarshal(&ctor.Param); err != nil {
		return nil, fmt.Errorf("%s: %w", annot.Decl.Name, err)
	}
	ctor.Annot = annot
	return &ctor, nil
}
