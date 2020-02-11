package typcore

import (
	log "github.com/sirupsen/logrus"
)

type appDescriptor interface {
	RunApp() error
}

type buildDescriptor interface {
	RunBuild() error
}

// RunApp the application
func RunApp(d appDescriptor) {
	if err := d.RunApp(); err != nil {
		log.Fatal(err.Error())
	}
}

// RunBuildTool the build tool
func RunBuildTool(d buildDescriptor) {
	if err := d.RunBuild(); err != nil {
		log.Fatal(err.Error())
	}
}
