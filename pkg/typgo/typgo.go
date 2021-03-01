package typgo

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

// BuildTool app
func BuildTool(d *Descriptor) *cli.App {

	if d.Environment != nil {
		envContext := NewPrepareContext(d, "load-env")
		if err := d.Environment.EnvLoad(envContext); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}

	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	for _, task := range d.Tasks {
		app.Commands = append(app.Commands, task.Task().CliCommand(d))
	}

	return app
}

// Start typical build-tool
func Start(d *Descriptor) {
	if err := BuildTool(d).Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
