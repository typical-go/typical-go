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
		var (
			f *os.File
		)
		if err = d.Validate(); err != nil {
			return
		}
		if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
			log.Infof("Generate new project environment at '%s'", defaultDotEnv)
			if f, err = os.Create(defaultDotEnv); err != nil {
				return
			}
			defer f.Close()
			_, configMap := typcore.CreateConfigMap(d)
			if err = WriteEnv(f, configMap); err != nil {
				return
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
	ctx := typcore.NewContext(d)
	for _, module := range d.AllModule() {
		if commander, ok := module.(typcore.BuildCommander); ok {
			for _, cmd := range commander.BuildCommands(ctx) {
				cmds = append(cmds, cmd)
			}
		}
	}
	return
}
