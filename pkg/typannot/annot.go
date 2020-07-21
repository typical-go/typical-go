package typannot

import (
	"encoding/json"
	"errors"
	"strings"
)

type (
	// Annot that contain extra additional information
	Annot struct {
		TagName  string `json:"tag_name"`
		TagAttrs []byte `json:"tag_attrs"`
		*Decl    `json:"decl"`
	}
)

// CreateAnnot parse raw string to annotation
func CreateAnnot(decl *Decl, raw string) (a *Annot, err error) {

	if !strings.HasPrefix(raw, "@") {
		return nil, errors.New("Annotation: should start with @")
	}
	raw = raw[1:]

	i1 := strings.IndexRune(raw, '{')
	if i1 < 0 {
		return &Annot{
			Decl:    decl,
			TagName: strings.TrimSpace(raw),
		}, nil
	}

	i2 := strings.IndexRune(raw, '}')
	if i2 < 0 {
		return nil, errors.New("Annotation: missing '}'")
	}

	return &Annot{
		Decl:     decl,
		TagName:  strings.TrimSpace(raw[:i1]),
		TagAttrs: []byte(strings.TrimSpace(raw[i1 : i2+1])),
	}, nil
}

// Unmarshal tag attributes
func (a *Annot) Unmarshal(v interface{}) error {
	if len(a.TagAttrs) > 0 {
		return json.Unmarshal(a.TagAttrs, v)
	}
	return nil
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
