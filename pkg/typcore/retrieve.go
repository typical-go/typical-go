package typcore

import (
	"fmt"
	"go/build"
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

	for _, source := range sources {
		if _, err = os.Stat(source); os.IsNotExist(err) {
			return nil, fmt.Errorf("ProjectSource '%s' is not exist", source)
		}
	}
	return
}

// RetrieveProjectPackage to retrieve project package
func RetrieveProjectPackage() (pkg string) {
	var (
		err  error
		root string
		f    *os.File
	)

	if root, err = os.Getwd(); err != nil {
		panic(err.Error())
	}

	if f, err = os.Open(root + "/go.mod"); err != nil {
		// NOTE: go.mod is not exist. Check if the project sit in $GOPATH
		gopath := build.Default.GOPATH
		if strings.HasPrefix(root, gopath) {
			return root[len(gopath):]
		}
		panic("Failed to retrieve ProjectPackage: `go.mod` is missing and the project not in $GOPATH")
	}
	defer f.Close()

	modfile := common.ParseModfile(f)
	return modfile.ProjectPackage
}
