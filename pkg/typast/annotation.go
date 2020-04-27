package typast

// Annotation contain extra additional information
type Annotation struct {
	TagName  string
	TagAttrs map[string]string
}

// NewAnnotation return new instance of Annotation
func NewAnnotation(tagName string) *Annotation {
	return &Annotation{
		TagName:  tagName,
		TagAttrs: make(map[string]string),
	}
}

// PutAttr to put attribute
func (a *Annotation) PutAttr(key, value string) *Annotation {
	a.TagAttrs[key] = value
	return a
}
