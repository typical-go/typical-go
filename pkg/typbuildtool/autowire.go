package typbuildtool

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

// Autowires is list of function declarion to be provided by dig
type Autowires []string

// OnAnnotation to handle annotation event
func (a *Autowires) OnAnnotation(e *walker.AnnotationEvent) (err error) {
	*a = append(*a, fmt.Sprintf("%s.%s", e.File.Name, e.SourceName))
	return
}
