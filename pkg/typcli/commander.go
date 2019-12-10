package typcli

import "github.com/urfave/cli/v2"

// BuildCommander responsible to command
type BuildCommander interface {
	BuildCommands(c *Container) []*cli.Command
}

// AppCommander return command
type AppCommander interface {
	AppCommands(c *Container) []*cli.Command
}

// IsBuildCommander return true if obj implement commander
func IsBuildCommander(obj interface{}) (ok bool) {
	_, ok = obj.(BuildCommander)
	return
}

// IsAppCommander return true if object implementation of AppCLI
func IsAppCommander(obj interface{}) (ok bool) {
	_, ok = obj.(AppCommander)
	return
}
