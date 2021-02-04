package typgo

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Cli app for typical-go
func Cli(b *BuildSys) *cli.App {

	if b.EnvLoader != nil {
		if err := b.EnvLoader.EnvLoad(); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
	}

	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	for _, task := range b.Tasks {
		app.Commands = append(app.Commands, task.Task(b))
	}

	return app
}
