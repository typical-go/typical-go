package typgo

import (
	"github.com/urfave/cli/v2"
)

// Cli app for typical-go
func Cli(b *BuildSys) *cli.App {
	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	app.Before = func(*cli.Context) error {
		if b.EnvLoader != nil {
			return b.EnvLoader.EnvLoad()
		}
		return nil
	}
	app.Commands = b.Commands
	return app
}
