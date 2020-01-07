package typbuildtool

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

const (
	defaultDotEnv = ".env"
)

// Run the build tool
func Run(d *typcore.ProjectDescriptor) {
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = ""
	app.Description = d.Description
	app.Version = d.Version
	app.Before = func(ctx *cli.Context) (err error) {
		var f *os.File
		if err = d.Validate(); err != nil {
			return
		}
		if d.Configuration != nil {
			if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
				log.Infof("Generate new project environment at '%s'", defaultDotEnv)
				if f, err = os.Create(defaultDotEnv); err != nil {
					return
				}
				defer f.Close()
				keys, configMap := d.Configuration.ConfigMap()
				if err = WriteEnv(f, keys, configMap); err != nil {
					return
				}
			}
		}
		return
	}
	app.Commands = commands(d)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}

func commands(d *typcore.ProjectDescriptor) (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdBuild(d),
		cmdClean(),
		cmdRun(d),
		cmdTest(),
		cmdMock(d),
	}
	if d.Releaser != nil {
		cmds = append(cmds, cmdRelease(d))
	}
	cmds = append(cmds, BuildCommands(d)...)
	return
}

// BuildCommands return list of command
func BuildCommands(d *typcore.ProjectDescriptor) (cmds []*cli.Command) {
	bc := typcore.NewBuildContext(d)
	for _, commander := range d.BuildCommands {
		for _, cmd := range commander.BuildCommands(bc) {
			cmds = append(cmds, cmd)
		}
	}
	return
}
