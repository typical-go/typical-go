package typcore

import "log"

// BuildToolLauncher responsible to launch the build-tool
type BuildToolLauncher interface {
	LaunchBuildTool() error
}

// LaunchBuildTool the build tool
func LaunchBuildTool(launcher BuildToolLauncher) {
	if err := launcher.LaunchBuildTool(); err != nil {
		log.Fatal(err.Error())
	}
}
