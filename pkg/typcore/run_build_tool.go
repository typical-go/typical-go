package typcore

import (
	log "github.com/sirupsen/logrus"
)

// RunBuildTool the build tool
func RunBuildTool(d *Descriptor) {
	var (
		err  error
		bctx *BuildContext
	)

	if bctx, err = d.BuildContext(); err != nil {
		log.Fatal(err.Error())
	}
	if d.Configuration != nil {
		if err := d.Configuration.Setup(); err != nil {
			log.Fatal(err.Error())
		}
	}
	if err = d.Build.Run(bctx); err != nil {
		log.Fatal(err.Error())
	}
}
