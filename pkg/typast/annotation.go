package typast

import (
	"strings"
)

// Annotation contain extra additional information
type Annotation struct {
	*Decl
	TagName  string
	TagAttrs map[string]string
}

// CreateAnnotation parse raw string to annotation
func CreateAnnotation(decl *Decl, raw string) (a *Annotation) {
	if !strings.HasPrefix(raw, "@") {
		panic("Annotation should start with @")
	}
	raw = raw[1:]
	return &Annotation{
		Decl:     decl,
		TagName:  tagName(raw),
		TagAttrs: tagAttrs(map[string]string{}, rawAttribute(raw)),
	}
}

// Equal return true if annotation have same tag name and declaration type
func (a *Annotation) Equal(tagName string, declType DeclType) bool {
	return strings.ToLower(tagName) == strings.ToLower(a.TagName) &&
		a.Type == declType
}

func tagName(raw string) string {
	i := strings.IndexRune(raw, '(')
	if i > 0 {
		raw = raw[:i]
	}
	return strings.TrimSpace(raw)
}

func rawAttribute(raw string) string {
	i0 := strings.IndexRune(raw, '(')
	i1 := strings.IndexRune(raw, ')')
	if i0 < 0 || i1 < 0 {
		return ""
	}
	return raw[i0+1 : i1]
}

func tagAttrs(m map[string]string, rawAttr string) map[string]string {
	var (
		key, value string
	)

	if m == nil {
		m = map[string]string{}
	}

	if rawAttr = strings.TrimSpace(rawAttr); rawAttr == "" {
		return m
	}

	eq := strings.IndexRune(rawAttr, '=')
	space := strings.IndexRune(rawAttr, ' ')

	if space > 0 && (space < eq || eq < 1) {
		key := strings.TrimSpace(rawAttr[:space])
		m[key] = ""
		tagAttrs(m, rawAttr[space+1:])
		return m
	}

	if eq < 0 {
		m[rawAttr] = ""
		return m
	}

	key = strings.TrimSpace(rawAttr[:eq])
	if eq == len(rawAttr)-1 {
		m[key] = ""
		return m
	}

	value = rawAttr[eq+1:]
	if value[0] == '"' {
		value = value[1:]
		i := strings.IndexRune(value, '"')
		m[key] = value[:i]
		tagAttrs(m, value[i+1:])
		return m
	}
	if value[0] == ' ' {
		m[key] = ""
		tagAttrs(m, value)
		return m
	}
	if i := strings.IndexRune(value, ' '); i > 0 {
		m[key] = value[:i]
		tagAttrs(m, value[i+1:])
		return m
	}
	m[key] = value
	return m
}
