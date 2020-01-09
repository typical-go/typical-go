package typbuildtool

import (
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

// Automocks is list of filename to be mocking by mockgen
type Automocks []string

// OnAnnotation handle type specificatio event
func (a *Automocks) OnAnnotation(e *walker.AnnotationEvent) (err error) {
	*a = append(*a, e.Filename)
	return
}
