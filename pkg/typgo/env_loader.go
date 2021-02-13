package typgo

import (
	"github.com/typical-go/typical-go/pkg/envkit"
)

type (
	// EnvLoader responsible to load env
	EnvLoader interface {
		EnvLoad(*BuildToolContext) error
	}
	// DotEnv file
	DotEnv string
	// EnvMap environment map
	EnvMap map[string]string
)

// DotEnv

var _ EnvLoader = (DotEnv)("")

// EnvLoad load environment from dotenv file
func (d DotEnv) EnvLoad(c *BuildToolContext) error {
	m, _ := envkit.ReadFile(string(d))
	if len(m) > 0 {
		c.Infof("Read from DotEnv '%s': %s\n", d, m.SortedKeys())
		return envkit.Setenv(m)
	}
	return nil
}

// EnvMap

var _ EnvLoader = (EnvMap)(nil)

// EnvLoad load environment from dotenv file
func (e EnvMap) EnvLoad(c *BuildToolContext) error {
	m := envkit.Map(e)
	c.Infof("Read from EnvMap: %s\n", m.SortedKeys())
	return envkit.Setenv(m)
}
