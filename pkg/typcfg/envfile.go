package typcfg

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

const configKey = "CONFIG"

// LoadEnvFile to load environment from .env file
func LoadEnvFile() (err error) {
	// TODO: don't use godotenv for flexibility
	configSource := os.Getenv(configKey)
	var configs []string
	var envMap map[string]string
	if configSource == "" {
		envMap, _ = godotenv.Read()
	} else {
		configs = strings.Split(configSource, ",")
		if envMap, err = godotenv.Read(configs...); err != nil {
			return
		}
	}

	if len(envMap) > 0 {
		log.Infof("Load environments %s", configSource)
		var b strings.Builder
		for key, value := range envMap {
			if err = os.Setenv(key, value); err != nil {
				return
			}
			b.WriteString("+")
			b.WriteString(key)
			b.WriteString(" ")
		}
		log.Info(b.String())
	}
	return
}
