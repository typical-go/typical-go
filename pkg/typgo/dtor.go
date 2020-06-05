package typgo

import (
	"github.com/typical-go/typical-go/pkg/typast"
)

var (
	dtorTags = []string{
		"dtor",
	}
)

// Dtor is destructor tag
type Dtor struct {
	*typast.Annot `json:"-"`
}

// CreateDtor to create create Dtor annotation
func CreateDtor(annot *typast.Annot) *Dtor {
	if !IsFuncTag(annot, dtorTags...) {
		return nil
	}

	return &Dtor{
		Annot: annot,
	}
}
