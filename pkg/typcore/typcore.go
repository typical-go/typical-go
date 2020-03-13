package typcore

import (
	"go/build"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/common"
)

// RetrieveProjectSources to retrieve project source from descriptor
func RetrieveProjectSources(d *Descriptor) (sources []string) {
	if sourceable, ok := d.App.(SourceableApp); ok {
		sources = append(sources, sourceable.ProjectSources()...)
	} else {
		sources = append(sources, common.PackageName(d.App))
	}
	if _, err := os.Stat("pkg"); !os.IsNotExist(err) {
		sources = append(sources, "pkg")
	}
	return
}

// RetrieveProjectPackage to retrieve module package from root path
func RetrieveProjectPackage(root string) (pkg string) {
	var (
		err error
	)

	var f *os.File
	if f, err = os.Open(root + "/go.mod"); err != nil {
		// NOTE: go.mod is not exist. Check if the project sit in $GOPATH
		gopath := build.Default.GOPATH
		if strings.HasPrefix(root, gopath) {
			return root[len(gopath):]
		}

		log.Warn("Can't get default module package. `go.mod` is missing and the project not in $GOPATH")
		return ""
	}
	defer f.Close()

	modfile := common.ParseModfile(f)
	return modfile.ProjectPackage
}
