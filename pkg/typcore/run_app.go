package typcore

import (
	log "github.com/sirupsen/logrus"
)

// RunApp the application
func RunApp(d *Descriptor) {
	var (
		actx *AppContext
		err  error
	)
	if actx, err = d.AppContext(); err != nil {
		log.Fatal(err.Error())
	}
	if err = d.App.Run(actx); err != nil {
		log.Fatal(err.Error())
	}
}
