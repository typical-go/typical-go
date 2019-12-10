package typbuildtool

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/typical-go/typical-go/pkg/typobj"
	"github.com/urfave/cli/v2"
)

// Run the build tool
func Run(c *typctx.Context) {
	buildtool := buildtool{Context: c}
	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = ""
	app.Description = c.Description
	app.Version = c.Version
	app.Before = buildtool.before
	app.Commands = buildtool.commands()
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}

// ModuleCommands return list of command
func ModuleCommands(ctx *typctx.Context) (cmds []*cli.Command) {
	for _, module := range ctx.AllModule() {
		buildCli := typctx.NewCli(ctx, module)
		if commander, ok := module.(typobj.BuildCommander); ok {
			for _, cmd := range commander.BuildCommands(buildCli) {
				cmds = append(cmds, cmd)
			}
		}
	}
	return
}

type buildtool struct {
	*typctx.Context
}

func (t buildtool) commands() (cmds []*cli.Command) {
	cmds = []*cli.Command{
		t.cmdBuild(),
		t.cmdClean(),
		t.cmdRun(),
		t.cmdTest(),
		t.cmdMock(),
	}
	if t.Releaser != nil {
		cmds = append(cmds, t.cmdRelease())
	}
	if t.ReadmeGenerator != nil {
		cmds = append(cmds, t.cmdReadme())
	}
	cmds = append(cmds, ModuleCommands(t.Context)...)
	return
}

func (t buildtool) before(ctx *cli.Context) (err error) {
	return t.Context.Validate()
}
