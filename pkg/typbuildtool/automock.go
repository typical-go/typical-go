package typbuildtool

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

// Automocks is list of filename to be mocking by mockgen
type Automocks []string

// OnDecl handle type specificatio event
func (a *Automocks) OnDecl(e *walker.DeclEvent) (err error) {
	annotations := e.Doc.Annotations()
	if annotations.Contain("mock") {
		if e.EventType == walker.InterfaceType {
			*a = append(*a, e.Filename)
		} else {
			log.Warnf("[mock] has no effect to %s:%s", e.EventType, e.Name)
		}
	}

	return
}
