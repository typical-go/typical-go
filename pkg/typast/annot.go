package typast

import (
	"encoding/json"
	"errors"
	"strings"
)

// Annot is annotation that contain extra additional information
type Annot struct {
	*Decl
	TagName  string
	TagAttrs []byte
}

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
