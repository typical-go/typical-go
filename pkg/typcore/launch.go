package typcore

import "log"

type (

	// BuildLauncher responsible to launch the build-tool
	BuildLauncher interface {
		LaunchBuild() error
	}
)

// LaunchBuild to launch the build tool
func LaunchBuild(launcher BuildLauncher) {
	if err := launcher.LaunchBuild(); err != nil {
		log.Fatal(err.Error())
	}
}
