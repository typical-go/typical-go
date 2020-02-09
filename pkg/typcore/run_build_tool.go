package typcore

import (
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	defaultDotEnv = ".env"
)

// RunBuildTool the build tool
func RunBuildTool(d *Descriptor) {
	var (
		f    *os.File
		err  error
		bctx *BuildContext
	)

	if bctx, err = d.BuildContext(); err != nil {
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
	if err = d.Build.Run(bctx); err != nil {
		log.Fatal(err.Error())
	}
}
