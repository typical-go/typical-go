package typgo

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Cli app for typical-go
func Cli(d *Descriptor) *cli.App {

	if d.EnvLoader != nil {
		if err := d.EnvLoader.EnvLoad(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}

	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	for _, task := range d.Tasks {
		app.Commands = append(app.Commands, task.Task(d))
	}

	return app
}
