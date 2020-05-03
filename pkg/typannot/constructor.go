package typannot

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"

	"github.com/typical-go/typical-go/pkg/typast"
)

var (
	// ConstructorTag is constructor tag
	ConstructorTag = "constructor"
)

// Constructor is contructor annotation
type Constructor struct {
	Name string `json:"name"`
	Def  string `json:"-"`
}

// GetConstructor to get contructor annotation data
func GetConstructor(store *typast.ASTStore) (annots []*Constructor, errs common.Errors) {
	var err error
	for _, annot := range store.Annots {
		if isConstructor(annot) {
			var ctorAnnot Constructor
			if len(annot.TagAttrs) > 0 {
				if err = json.Unmarshal(annot.TagAttrs, &ctorAnnot); err != nil {
					errs.Append(fmt.Errorf("Invalid tag attribute %s", annot.TagAttrs))
					continue
				}
			}
			ctorAnnot.Def = fmt.Sprintf("%s.%s", annot.Pkg, annot.Name)
			annots = append(annots, &ctorAnnot)
		}
	}
	return
}

func isConstructor(annot *typast.Annotation) bool {
	return strings.EqualFold(annot.TagName, ConstructorTag) &&
		annot.Type == typast.Function
}
