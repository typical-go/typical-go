package typcore

import (
	"os"

	"github.com/typical-go/typical-go/pkg/common"
)

var (
	// DefaultModulePackage is default value for ModulePackage
	DefaultModulePackage = "" // NOTE: supply by ldflags

	// DefaultTempFolder is default value for temp folder location
	DefaultTempFolder = ".typical-tmp"

	// DefaultCmdFolder is default value for cmd folder location
	DefaultCmdFolder = "cmd"

	// DefaultBinFolder is default value for bin folder location
	DefaultBinFolder = "bin"

	// DefaultReleaseFolder is default value for release folder location
	DefaultReleaseFolder = "release"
)

// DefaultProjectSources to determine default project source
func DefaultProjectSources(d *Descriptor) (sources []string) {
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
