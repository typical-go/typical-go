package typcore

import (
	"go/build"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/common"
)

const (
	// Version of Typical-Go
	Version = "0.9.37"
)

var (
	// DefaultProjectPackage is default value for ProjectPackage
	DefaultProjectPackage = "" // NOTE: supply by ldflags

	// DefaultTempFolder is default value for temp folder location
	DefaultTempFolder = ".typical-tmp"

	// DefaultCmdFolder is default value for cmd folder location
	DefaultCmdFolder = "cmd"

	// DefaultBinFolder is default value for bin folder location
	DefaultBinFolder = "bin"

	// DefaultReleaseFolder is default value for release folder location
	DefaultReleaseFolder = "release"
)

// RetrieveProjectSources to retrieve project source from descriptor
func RetrieveProjectSources(d *Descriptor) (sources []string) {
	if sourceable, ok := d.App.(Sourceable); ok {
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

	gomod := buildkit.ParseGoMod(f)
	return gomod.ProjectPackage
}
