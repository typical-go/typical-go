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
	keys, configMap := c.ConfigMap()
	for _, key := range keys {
		var (
			v         interface{}
			cfgDetail = configMap[key]
		)
		if cfgDetail.IsZero {
			v = cfgDetail.Default
		} else {
			v = cfgDetail.Value
		}
		if _, err = fmt.Fprintf(w, "%s=%v\n", cfgDetail.Name, v); err != nil {
			return
		}
	}
	return
}
