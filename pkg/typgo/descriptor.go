package typgo

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type (
	// Descriptor describe the project
	Descriptor struct {
		ProjectName    string // By default is same with project folder. Only allowed characters(a-z,A-Z), underscore or dash.
		ProjectVersion string // By default it is 0.0.1
		EnvLoader      EnvLoader
		Tasks          []Tasker
	}
)

// BuildTool app
func BuildTool(d *Descriptor) *cli.App {
	if d.EnvLoader != nil {
		if err := d.EnvLoader.EnvLoad(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}

	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	for _, task := range d.Tasks {
		app.Commands = append(app.Commands, CliCommand(d, task.Task()))
	}

	return app
}

// Start typical build-tool
func Start(d *Descriptor) {
	if err := BuildTool(d).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
