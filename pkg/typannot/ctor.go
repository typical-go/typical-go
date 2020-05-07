package typannot

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typast"
)

var (
	// CtorTags is constructor tag
	CtorTags = []string{
		"ctor",
		"constructor",
	}
)

// Ctor is contructor annotation
type Ctor struct {
	*typast.Annot
	Name string `json:"name"`
}

// GetCtors to get contructor annotation data
func GetCtors(store *typast.ASTStore) (ctors []*Ctor, errs common.Errors) {
	for _, annot := range store.Annots {
		if IsFuncTag(annot, CtorTags) {
			ctor := new(Ctor)
			if err := annot.Unmarshal(ctor); err != nil {
				errs.Append(fmt.Errorf("%s: %w", CtorTags[0], err))
				continue
			}
			ctor.Annot = annot
			ctors = append(ctors, ctor)
		}
	}
	return
}
