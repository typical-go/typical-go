package walker

// Annotation contain extra additional information
type Annotation struct {
	Name string
}

// ParseAnnotation parse raw string to annotation
func ParseAnnotation(raw string) *Annotation {
	stripped := raw[1 : len(raw)-1]
	// TODO: handle parameter. Annotation Format: [name{key1=value1 key2=value2}]
	return &Annotation{
		Name: stripped,
	}
}
