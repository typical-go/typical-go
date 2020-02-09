package typcore

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/urfave/cli/v2"
)

const (
	defaultDotEnv = ".env"
)

// RunBuildTool the build tool
func RunBuildTool(d *Descriptor) {
	var (
		f   *os.File
		err error
		bc  *BuildContext
	)

	if bc, err = d.BuildContext(); err != nil {
		log.Fatal(err.Error())
	}
	if d.Configuration != nil {
		if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
			log.Infof("Generate new project environment at '%s'", defaultDotEnv)
			if f, err = os.Create(defaultDotEnv); err != nil {
				log.Fatal(err.Error())
			}
			defer f.Close()
			if err = d.Configuration.Write(f); err != nil {
				log.Fatal(err.Error())
			}
		}
	}

	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = d.Description
	app.Version = d.Version
	app.Commands = bc.Build.BuildCommands(bc)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
