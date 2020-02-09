package typcfg

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore"
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

// ConfigMap return map of config detail
func (c *Configuration) ConfigMap() (keys []string, configMap typcore.ConfigMap) {
	configMap = make(map[string]typcore.ConfigDetail)
	for _, configurer := range c.configurers {
		prefix, spec, _ := configurer.Configure(c.loader)
		details := typcore.CreateConfigDetails(prefix, spec)
		for _, detail := range details {
			name := detail.Name
			configMap[name] = detail
			keys = append(keys, name)
		}
	}
	return
}
