package typbuildtool

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/urfave/cli"
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
func ModuleCommands(ctx *typctx.Context) (cmds []cli.Command) {
	for _, module := range ctx.AllModule() {
		if cmd := Command(ctx, module); cmd != nil {
			cmds = append(cmds, *cmd)
		}
	}
	return
}

// Command of module
func Command(ctx *typctx.Context, obj interface{}) *cli.Command {
	if commander, ok := obj.(typcli.BuildCommander); ok {
		cmd := commander.BuildCommand(&typcli.ContextCli{
			Context: ctx,
		})
		return &cmd
	}
	if commander, ok := obj.(typcli.Commander); ok {
		cmd := commander.Command(&typcli.Cli{
			Obj:          obj,
			ConfigLoader: ctx.ConfigLoader,
		})
		return &cmd
	}
	return nil
}

type buildtool struct {
	*typctx.Context
}

func (t buildtool) commands() (cmds []cli.Command) {
	cmds = []cli.Command{
		t.cmdBuild(),
		t.cmdClean(),
		t.cmdRun(),
		t.cmdTest(),
		t.cmdRelease(),
		t.cmdMock(),
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
