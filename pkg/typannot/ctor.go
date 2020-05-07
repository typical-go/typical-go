package typannot

import (
	"encoding/json"
	"fmt"
	"strings"

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
	Name string `json:"name"`
	Def  string `json:"-"`
}

// GetCtor to get contructor annotation data
func GetCtor(store *typast.ASTStore) (annots []*Ctor, errs common.Errors) {
	var err error
	for _, a := range store.Annots {
		if isCtor(a) {
			var ctorAnnot Ctor
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

func isCtor(annot *typast.Annotation) bool {
	if annot.Type == typast.Function {
		for _, ctorTag := range CtorTags {
			if strings.EqualFold(ctorTag, annot.TagName) {
				return true
			}
		}
	}

	return false
}
