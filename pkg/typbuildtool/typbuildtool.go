package typbuildtool

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

// Run the build tool
func Run(c *typcore.Context) {
	var filenames []string
	var dirs []string
	var err error
	if dirs, filenames, err = projectFiles(typenv.Layout.App); err != nil {
		log.Fatal(err.Error())
	}
	buildtool := buildtool{
		Context:   c,
		filenames: filenames,
		dirs:      dirs,
	}
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

type buildtool struct {
	*typcore.Context
	filenames []string
	dirs      []string
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
	cmds = append(cmds, BuildCommands(t.Context)...)
	return
}

// BuildCommands return list of command
func BuildCommands(ctx *typcore.Context) (cmds []*cli.Command) {
	for _, module := range ctx.AllModule() {
		buildCli := typcore.NewCli(ctx, module)
		if commander, ok := module.(typcore.BuildCommander); ok {
			for _, cmd := range commander.BuildCommands(buildCli) {
				cmds = append(cmds, cmd)
			}
		}
	}
	return
}
