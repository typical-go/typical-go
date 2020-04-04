package typcfg

import "github.com/kelseyhightower/envconfig"

// Process populates the specified struct based on environment variables
func Process(name string, spec interface{}) error {
	return envconfig.Process(name, spec)
}
