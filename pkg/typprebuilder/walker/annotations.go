package walker

import (
	"regexp"
	"strings"
)

// Annotations is list of annotation
type Annotations []*Annotation

// ParseAnnotations to parse godoc comment to list of annotation
func ParseAnnotations(doc string) (n Annotations) {
	r, _ := regexp.Compile("\\[(.*?)\\]")
	for _, s := range r.FindAllString(doc, -1) {
		n = append(n, ParseAnnotation(s))
	}
	return
}

// Get to get annotation by name
func (a Annotations) Get(name string) *Annotation {
	for _, notation := range a {
		if strings.ToLower(name) == strings.ToLower(notation.Name) {
			return notation
		}
	}
	return nil
}

// Contain return true if annotation exist
func (a Annotations) Contain(name string) bool {
	return a.Get(name) != nil
}
