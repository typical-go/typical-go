package common

import (
	"time"

	log "github.com/sirupsen/logrus"
)

// ElapsedTimeFn to print elapsed time of function
func ElapsedTimeFn(name string) func() {
	start := time.Now()
	return func() {
		log.Debugf("%s took %v\n", name, time.Since(start))
	}
}
