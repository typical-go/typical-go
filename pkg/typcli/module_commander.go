package typcli

import "github.com/urfave/cli/v2"

// ModuleCommander responsible to command
type ModuleCommander interface {
	Command(c *ModuleCli) *cli.Command
}

// IsModuleCommander return true if obj implement commander
func IsModuleCommander(obj interface{}) (ok bool) {
	_, ok = obj.(ModuleCommander)
	return
}
