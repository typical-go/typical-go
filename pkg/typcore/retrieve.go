package typcore

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/typical-go/typical-go/pkg/common"
)

// RetrievePackageName return package name of the interface
func RetrievePackageName(v interface{}) string {
	if common.IsNil(v) {
		return ""
	}

	s := reflect.TypeOf(v).String()
	if dot := strings.Index(s, "."); dot > 0 {
		if strings.HasPrefix(s, "*") {
			return s[1:dot]
		}
		return s[:dot]
	}
	return ""
}

// RetrieveProjectSources to retrieve project source
func RetrieveProjectSources(d *Descriptor) (sources []string, err error) {
	if sourceable, ok := d.App.(SourceableApp); ok {
		sources = append(sources, sourceable.ProjectSources()...)
	} else {
		sources = append(sources, RetrievePackageName(d.App))
	}
	if _, err := os.Stat("pkg"); !os.IsNotExist(err) {
		sources = append(sources, "pkg")
	}
	if err = validateProjectSources(sources); err != nil {
		return nil, err
	}
	return
}

func validateProjectSources(sources []string) (err error) {
	for _, source := range sources {
		if _, err = os.Stat(source); os.IsNotExist(err) {
			return fmt.Errorf("ProjectSource '%s' is not exist", source)
		}
	}
	return
}
