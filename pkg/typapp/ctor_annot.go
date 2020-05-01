package typapp

import (
	"encoding/json"
	"fmt"

	"github.com/typical-go/typical-go/pkg/typbuildtool"

	"github.com/typical-go/typical-go/pkg/typast"
)

const (
	ctorAnnotTag = "constructor"
)

// CtorAnnot is contructor annotation
type CtorAnnot struct {
	Name string `json:"name"`
	Def  string `json:"-"`
}

// GetCtorAnnot to get c
func GetCtorAnnot(c *typbuildtool.PreconditionContext) (annots []*CtorAnnot) {
	var err error
	for _, a := range c.ASTStore().Annots {
		if a.Equal(ctorAnnotTag, typast.Function) {
			var ctorAnnot CtorAnnot
			if len(a.TagAttrs) > 0 {
				if err = json.Unmarshal(a.TagAttrs, &ctorAnnot); err != nil {
					c.Warnf("CtorAnnot: Invalid tag attribute %s", a.TagAttrs)
					continue
				}
			}
			ctorAnnot.Def = fmt.Sprintf("%s.%s", a.Pkg, a.Name)

			annots = append(annots, &ctorAnnot)
		}
	}
	return
}
