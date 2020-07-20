package typgo

import (
	"github.com/urfave/cli/v2"
)

type (
	// BuildSys detail
	BuildSys struct {
		*Descriptor
		Commands []*cli.Command
	}
)

func createBuildSys(d *Descriptor) *BuildSys {
	sys := &BuildSys{
		Descriptor: d,
	}
	for _, cmd := range d.Cmds {
		sys.Commands = append(sys.Commands, cmd.Command(sys))
	}
	return sys
}

func (b *BuildSys) app() *cli.App {
	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	app.Commands = b.Commands
	return app
}

// Run command by name
func (b *BuildSys) Run(name string, c *cli.Context) error {
	for _, command := range b.Commands {
		if command.Name == name {
			return command.Action(c)
		}
	}
	return nil
}

// ActionFn to return related action func
func (b *BuildSys) ActionFn(fn ExecuteFn) func(*cli.Context) error {
	return func(cliCtx *cli.Context) error {
		return fn(&Context{
			Context:  cliCtx,
			BuildSys: b,
		})
	}
}

// Execute action
func (b *BuildSys) Execute(action Action, cliCtx *cli.Context) error {
	return action.Execute(&Context{
		Context:  cliCtx,
		BuildSys: b,
	})
}
