package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/envkit"
	"github.com/urfave/cli/v2"
)

// Cli app for typical-go
func Cli(b *BuildSys) *cli.App {
	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	app.Before = beforeCliApp
	app.Commands = b.Commands
	return app
}

func beforeCliApp(*cli.Context) error {
	dotenv := ".env"
	m, _ := envkit.ReadFile(dotenv)
	if len(m) > 0 {
		fmt.Fprintf(Stdout, "Load environment '%s' %s\n", dotenv, m.SortedKeys())
		return envkit.Setenv(m)
	}
	return nil
}
