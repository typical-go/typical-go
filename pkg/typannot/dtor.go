package typannot

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typast"
)

var (
	dtorTags = []string{
		"dtor",
		"destructor",
	}
)

// Dtor is destructor tag
type Dtor struct {
	*typast.Annot
}

// GetDtors return dtor tag
func GetDtors(store *typast.ASTStore) (dtors []*Dtor, errs common.Errors) {
	for _, annot := range store.Annots {
		if IsFuncTag(annot, dtorTags) {
			dtor := new(Dtor)
			if err := annot.Unmarshal(dtor); err != nil {
				errs.Append(fmt.Errorf("%s: %w", dtorTags[0], err))
				continue
			}
			dtor.Annot = annot
			dtors = append(dtors, dtor)
		}
	}
	return
}
