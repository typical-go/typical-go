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

// Context of build-cli
func (b *BuildSys) Context(c *cli.Context) *Context {
	return &Context{
		Context:  c,
		BuildSys: b,
	}
}

// ActionFn to return related action func
func (b *BuildSys) ActionFn(fn ExecuteFn) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		return fn(b.Context(cli))
	}
}
