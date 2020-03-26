package typcore

import "log"

// AppLauncher responsible to launch the application
type AppLauncher interface {
	LaunchApp() error
}

// LaunchApp the application
func LaunchApp(launcher AppLauncher) {
	if err := launcher.LaunchApp(); err != nil {
		log.Fatal(err.Error())
	}
}
