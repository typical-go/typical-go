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
	sys := &BuildSys{Descriptor: d}
	for _, cmd := range d.Cmds {
		sys.Commands = append(sys.Commands, cmd.Command(sys))
	}
	return sys
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
