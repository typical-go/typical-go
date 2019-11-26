package envfile

import (
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

const configKey = "CONFIG"

// Load to load environment from .env file
func Load() (err error) {
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
	var b strings.Builder
	if len(envMap) > 0 {
		b.WriteString(fmt.Sprintf("Load environments %s: ", configSource))
		for key, value := range envMap {
			if err = os.Setenv(key, value); err != nil {
				return
			}
			b.WriteString(key)
			b.WriteString("; ")
		}
		log.Info(b.String())
	}
	return
}
