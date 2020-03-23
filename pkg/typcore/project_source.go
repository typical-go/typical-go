package typcore

import (
	"fmt"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
)

// ProjectSources of the project
func ProjectSources(d *Descriptor) (sources []string, err error) {
	if sourceable, ok := d.App.(SourceableApp); ok {
		sources = append(sources, sourceable.ProjectSources()...)
	} else {
		sources = append(sources, common.PackageName(d.App))
	}
	if _, err := os.Stat("pkg"); !os.IsNotExist(err) {
		sources = append(sources, "pkg")
	}

	for _, source := range sources {
		if _, err = os.Stat(source); os.IsNotExist(err) {
			return nil, fmt.Errorf("ProjectSource '%s' is not exist", source)
		}
	}
	return
}
