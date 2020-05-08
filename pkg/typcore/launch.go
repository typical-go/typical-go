package typcore

import "log"

type (
	// AppLauncher responsible to launch the application
	AppLauncher interface {
		LaunchApp() error
	}

	// BuildToolLauncher responsible to launch the build-tool
	BuildToolLauncher interface {
		LaunchBuildTool() error
	}
)

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
