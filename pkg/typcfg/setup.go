package typcfg

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

const (
	defaultDotEnv = ".env"
)

// Setup the configuration to be ready to use for the app and build-tool
func (c *Configuration) Setup() (err error) {
	var (
		f *os.File
	)
	if _, err = os.Stat(defaultDotEnv); os.IsNotExist(err) {
		log.Infof("Generate new project environment at '%s'", defaultDotEnv)
		if f, err = os.Create(defaultDotEnv); err != nil {
			return
		}
		defer f.Close()
		if err = c.Write(f); err != nil {
			return
		}
	}
	// TODO: load env
	return
}

func (c *Configuration) Write(w io.Writer) (err error) {
	store := c.Store()
	for _, field := range store.Fields(store.Keys()...) {
		var v interface{}
		if field.IsZero {
			v = field.Default
		} else {
			v = field.Value
		}
		if _, err = fmt.Fprintf(w, "%s=%v\n", field.Name, v); err != nil {
			return
		}
	}
	return
}
