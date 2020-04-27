package typast

// Annot stand of annotation that contain extra additional information
type Annot struct {
	Name  string
	Attrs map[string]string
}

// NewAnnot return new instance of Annot
func NewAnnot(name string) *Annot {
	return &Annot{
		Name:  name,
		Attrs: make(map[string]string),
	}
}

// Put attribute
func (a *Annot) Put(key, value string) *Annot {
	a.Attrs[key] = value
	return a
}
