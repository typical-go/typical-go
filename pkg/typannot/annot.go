package typannot

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

// CheckFunc return true if annot is function type with same tagName
func (a *Annot) CheckFunc(tagName string) bool {
	_, ok := a.Type.(*FuncType)
	return strings.EqualFold(tagName, a.TagName) && ok
}

// CheckInterface return true if annot is interface type with same tagName
func (a *Annot) CheckInterface(tagName string) bool {
	_, ok := a.Type.(*InterfaceType)
	return strings.EqualFold(tagName, a.TagName) && ok
}

// CheckStruct return true if annot is struct type with same tagName
func (a *Annot) CheckStruct(tagName string) bool {
	_, ok := a.Type.(*StructType)
	return strings.EqualFold(tagName, a.TagName) && ok
}
