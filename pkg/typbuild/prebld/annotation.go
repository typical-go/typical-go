package prebld

// Annotation contain extra additional information
type Annotation struct {
	Name  string
	Attrs map[string]string
}

// NewAnnotation return new instance of Annotation
func NewAnnotation(name string) *Annotation {
	return &Annotation{
		Name:  name,
		Attrs: make(map[string]string),
	}
}

// Put attribute
func (a *Annotation) Put(key, value string) *Annotation {
	a.Attrs[key] = value
	return a
}
