package typcore

import (
	log "github.com/sirupsen/logrus"
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
