package typbuildtool

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

// Constructors is list of function declarion to be provided by dig
type Constructors []string

// OnAnnotation to handle annotation event
func (a *Constructors) OnAnnotation(e *walker.AnnotationEvent) (err error) {
	*a = append(*a, fmt.Sprintf("%s.%s", e.File.Name, e.SourceName))
	return
}
