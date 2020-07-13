package typgo

import (
	"os"

	"github.com/urfave/cli/v2"
)

// Descriptor describe the project
type Descriptor struct {
	// Name of the project (OPTIONAL). It should be a characters with/without underscore or dash.
	// By default, project name is same with project folder
	Name string
	// Description of the project (OPTIONAL).
	Description string
	// Version of the project (OPTIONAL). By default it is 0.0.1
	Version  string
	Layouts  []string
	Commands Commands
}

// Run typical build-tool
func Run(d *Descriptor) error {
	buildCli := createBuildCli(d)

	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	app.Commands = buildCli.commands()

	return app.Run(os.Args)
}
