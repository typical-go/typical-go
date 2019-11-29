package typcli

import "github.com/urfave/cli/v2"

// AppCommander return command
type AppCommander interface {
	AppCommands(c Cli) []*cli.Command
}

// IsAppCommander return true if object implementation of AppCLI
func IsAppCommander(obj interface{}) (ok bool) {
	_, ok = obj.(AppCommander)
	return
}
