package typast

import (
	"reflect"
	"strings"
)

type (
	// Annot that contain extra additional information
	Annot struct {
		TagName  string            `json:"tag_name"`
		TagParam reflect.StructTag `json:"tag_param"`
		*Decl    `json:"decl"`
	}
)

func retrieveAnnots(decl *Decl) []*Annot {
	var annots []*Annot
	for _, raw := range decl.GetDocs() {
		if strings.HasPrefix(raw, "//") {
			raw = strings.TrimSpace(raw[2:])
		}
		if strings.HasPrefix(raw, "@") {
			tagName, tagAttrs := ParseAnnot(raw)
			annots = append(annots, &Annot{
				TagName:  tagName,
				TagParam: reflect.StructTag(tagAttrs),
				Decl:     decl,
			})
		}
	}

	return annots
}

// ParseAnnot parse raw string to annotation
func ParseAnnot(raw string) (tagName, tagAttrs string) {
	iOpen := strings.IndexRune(raw, '(')
	iSpace := strings.IndexRune(raw, ' ')

	if iOpen < 0 {
		if iSpace < 0 {
			tagName = strings.TrimSpace(raw)
			return tagName, ""
		}
		tagName = raw[:iSpace]
	} else {
		if iSpace < 0 {
			tagName = raw[:iOpen]
		} else {
			tagName = raw[:iSpace]
		}

		if iClose := strings.IndexRune(raw, ')'); iClose > 0 {
			tagAttrs = raw[iOpen+1 : iClose]
		}
	}

	return tagName, tagAttrs
}
