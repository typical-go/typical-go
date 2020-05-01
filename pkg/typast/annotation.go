package typast

import (
	"errors"
	"strings"
)

// Annotation contain extra additional information
type Annotation struct {
	*Decl
	TagName  string
	TagAttrs []byte
}

// CreateAnnotation parse raw string to annotation
func CreateAnnotation(decl *Decl, raw string) (a *Annotation, err error) {

	if !strings.HasPrefix(raw, "@") {
		return nil, errors.New("Annotation: should start with @")
	}
	raw = raw[1:]

	i1 := strings.IndexRune(raw, '{')
	if i1 < 0 {
		return &Annotation{
			Decl:    decl,
			TagName: strings.TrimSpace(raw),
		}, nil
	}

	i2 := strings.IndexRune(raw, '}')
	if i2 < 0 {
		return nil, errors.New("Annotation: missing '}'")
	}

	return &Annotation{
		Decl:     decl,
		TagName:  strings.TrimSpace(raw[:i1]),
		TagAttrs: []byte(strings.TrimSpace(raw[i1 : i2+1])),
	}, nil
}

// Equal return true if annotation have same tag name and declaration type
func (a *Annotation) Equal(tagName string, declType DeclType) bool {
	return strings.ToLower(tagName) == strings.ToLower(a.TagName) &&
		a.Type == declType
}
