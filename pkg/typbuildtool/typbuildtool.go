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
	var (
		f   *os.File
		err error
		bc  *typcore.BuildContext
	)
	if err = d.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	if bc, err = typcore.CreateBuildContext(d); err != nil {
		log.Fatal(err.Error())
	}
	if d.Configuration != nil {
		if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
			log.Infof("Generate new project environment at '%s'", defaultDotEnv)
			if f, err = os.Create(defaultDotEnv); err != nil {
				log.Fatal(err.Error())
			}
			defer f.Close()
			keys, configMap := d.Configuration.ConfigMap()
			if err = WriteEnv(f, keys, configMap); err != nil {
				log.Fatal(err.Error())
			}
		}
	}

	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version
	app.Commands = BuildCommands(bc)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}

// BuildCommands return list of command
func BuildCommands(bc *typcore.BuildContext) (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdBuild(bc),
		cmdClean(),
		cmdRun(bc),
		cmdTest(),
		cmdMock(bc),
	}
	if bc.Releaser != nil {
		cmds = append(cmds, cmdRelease(bc))
	}
	for _, commander := range bc.BuildCommands {
		for _, cmd := range commander.BuildCommands(bc) {
			cmds = append(cmds, cmd)
		}
	}
	return
}
