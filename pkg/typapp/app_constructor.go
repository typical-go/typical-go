package typapp

import "github.com/typical-go/typical-go/pkg/typdep"

var (
	appConstructors []*typdep.Constructor
)

// AppendConstructor to append constructor
func AppendConstructor(cons ...*typdep.Constructor) {
	appConstructors = append(appConstructors, cons...)
}
