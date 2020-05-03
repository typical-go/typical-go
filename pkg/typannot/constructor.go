package typannot

import (
	"encoding/json"
	"fmt"

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
	for _, a := range store.Annots {
		if a.Equal(ConstructorTag, typast.Function) {
			var ctorAnnot Constructor
			if len(a.TagAttrs) > 0 {
				if err = json.Unmarshal(a.TagAttrs, &ctorAnnot); err != nil {
					errs.Append(fmt.Errorf("Invalid tag attribute %s", a.TagAttrs))
					continue
				}
			}
			ctorAnnot.Def = fmt.Sprintf("%s.%s", a.Pkg, a.Name)
			annots = append(annots, &ctorAnnot)
		}
	}
	return
}
