package typcli

import "github.com/urfave/cli/v2"

// BuildCommander responsible to command
type BuildCommander interface {
	BuildCommands(c *BuildCli) []*cli.Command
}

// IsBuildCommander return true if obj implement commander
func IsBuildCommander(obj interface{}) (ok bool) {
	_, ok = obj.(BuildCommander)
	return
}
