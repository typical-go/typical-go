package typcore

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/common"
)

var (
	DefaultModulePackage string
)

// RunApp the application
func RunApp(i interface{ RunApp() error }) {
	if err := i.RunApp(); err != nil {
		log.Fatal(err.Error())
	}
}

// RunBuildTool the build tool
func RunBuildTool(i interface{ RunBuild() error }) {
	if err := i.RunBuild(); err != nil {
		log.Fatal(err.Error())
	}
}

// DefaultProjectSources return default project source
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
