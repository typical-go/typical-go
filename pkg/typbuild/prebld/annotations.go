package prebld

import (
	"strings"
)

// Annotations is list of annotation
type Annotations []*Annotation

// Get to get annotation by name
func (a Annotations) Get(name string) *Annotation {
	for _, a := range a {
		if strings.ToLower(name) == strings.ToLower(a.Name) {
			return a
		}
	}
	return nil
}

// Contain return true if annotation exist
func (a Annotations) Contain(name string) bool {
	return a.Get(name) != nil
}
