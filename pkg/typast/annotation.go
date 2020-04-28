package typast

import "strings"

// Annotation contain extra additional information
type Annotation struct {
	TagName  string
	TagAttrs map[string]string
}

// CreateAnnotation parse raw string to annotation
func CreateAnnotation(raw string) (a *Annotation) {
	if raw[0] != '[' && raw[len(raw)-1] != ']' {
		return
	}
	raw = raw[1 : len(raw)-1]

	return &Annotation{
		TagName:  tagName(raw),
		TagAttrs: tagAttrs(map[string]string{}, rawAttribute(raw)),
	}
}

func tagName(raw string) string {
	i := strings.IndexRune(raw, '(')
	if i < 0 {
		return raw
	}
	return raw[:i]
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
		m[rawAttr[:space]] = ""
		tagAttrs(m, rawAttr[space+1:])
		return m
	}
	if eq > 0 {
		key = rawAttr[:eq]
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

	m[rawAttr] = ""
	return m
}

// PutAttr to put attribute
func (a *Annotation) PutAttr(key, value string) *Annotation {
	a.TagAttrs[key] = value
	return a
}
