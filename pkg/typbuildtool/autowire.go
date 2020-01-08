package typbuildtool

import (
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

// Autowires is list of function declarion to be provided by dig
type Autowires []string

// OnDecl to handle declaration event
func (a *Autowires) OnDecl(e *walker.DeclEvent) (err error) {
	annotations := e.Doc.Annotations()
	if annotations.Contain("autowire") {
		if e.EventType == walker.FunctionType {
			*a = append(*a, fmt.Sprintf("%s.%s", e.File.Name, e.Name))
		} else {
			log.Warnf("[autowire] has no effect to %s:%s", e.EventType, e.Name)
		}
	}
	return
}
