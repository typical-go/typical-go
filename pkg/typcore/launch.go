package typcore

import "log"

type (
	// AppLauncher responsible to launch the application
	AppLauncher interface {
		LaunchApp() error
	}

	// BuildLauncher responsible to launch the build-tool
	BuildLauncher interface {
		LaunchBuild() error
	}
)

// LaunchApp to launch the application
func LaunchApp(launcher AppLauncher) {
	if err := launcher.LaunchApp(); err != nil {
		log.Fatal(err.Error())
	}
}

// LaunchBuild to launch the build tool
func LaunchBuild(launcher BuildLauncher) {
	if err := launcher.LaunchBuild(); err != nil {
		log.Fatal(err.Error())
	}
}
