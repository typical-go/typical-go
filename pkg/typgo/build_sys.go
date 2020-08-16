package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
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
	sys := &BuildSys{Descriptor: d}
	for _, cmd := range d.Cmds {
		sys.Commands = append(sys.Commands, cmd.Command(sys))
	}
	return sys
}

func (b *BuildSys) app() *cli.App {
	cli.AppHelpTemplate = appHelpTemplate
	cli.SubcommandHelpTemplate = subcommandHelpTemplate

	app := cli.NewApp()
	app.Before = func(*cli.Context) error {
		return common.LoadEnv(".env")
	}
	app.Commands = b.Commands
	return app
}

// Run command by name
func (b *BuildSys) Run(name string, c *cli.Context) error {
	for _, command := range b.Commands {
		if command.Name == name {
			if command.Before != nil {
				if err := command.Before(c); err != nil {
					return err
				}
			}
			return command.Action(c)
		}
	}
	return fmt.Errorf("typgo: no command with name '%s'", name)
}

// ExecuteFn to return cli func from execute function
func (b *BuildSys) ExecuteFn(fn ExecuteFn) func(*cli.Context) error {
	return b.Action(NewAction(fn))
}

// Action to return cli func from action
func (b *BuildSys) Action(action Action) func(*cli.Context) error {
	return func(cliCtx *cli.Context) error {
		if action == nil {
			return nil
		}
		return action.Execute(&Context{
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
