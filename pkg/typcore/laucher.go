package typcore

import (
	log "github.com/sirupsen/logrus"
)

// AppLauncher responsible to launch the application
type AppLauncher interface {
	LaunchApp() error
}

// BuildToolLauncher responsible to launch the build-tool
type BuildToolLauncher interface {
	LaunchBuildTool() error
}

// LaunchApp the application
func LaunchApp(launcher AppLauncher) {
	if err := launcher.LaunchApp(); err != nil {
		log.Fatal(err.Error())
	}
}

// LaunchBuildTool the build tool
func LaunchBuildTool(launcher BuildToolLauncher) {
	if err := launcher.LaunchBuildTool(); err != nil {
		log.Fatal(err.Error())
	}
}
