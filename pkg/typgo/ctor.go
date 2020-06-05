package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typast"
)

var (
	ctorTags = []string{
		"ctor",
	}
)

type (
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

// CreateCtor create ctor annotation
func CreateCtor(annot *typast.Annot) (*Ctor, error) {
	if !IsFuncTag(annot, ctorTags...) {
		return nil, nil
	}

	ctor := new(Ctor)
	if err := annot.Unmarshal(&ctor.Param); err != nil {
		return nil, fmt.Errorf("%s: %w", annot.Decl.Name, err)
	}
	ctor.Annot = annot
	return ctor, nil
}
