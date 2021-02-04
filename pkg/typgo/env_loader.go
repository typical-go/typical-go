package typgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/envkit"
	"github.com/typical-go/typical-go/pkg/oskit"
)

type (
	// EnvLoader responsible to load env
	EnvLoader interface {
		EnvLoad() error
	}
	// DotEnv file
	DotEnv string
	// EnvMap environment map
	EnvMap map[string]string
)

// DotEnv

var _ EnvLoader = (DotEnv)("")

// EnvLoad load environment from dotenv file
func (d DotEnv) EnvLoad() error {
	m, _ := envkit.ReadFile(string(d))
	if len(m) > 0 {
		fmt.Fprintf(oskit.Stdout, "Load environment from '%s': %s\n\n", d, m.SortedKeys())
		return envkit.Setenv(m)
	}
	return nil
}

// EnvMap

var _ EnvLoader = (EnvMap)(nil)

// EnvLoad load environment from dotenv file
func (e EnvMap) EnvLoad() error {
	m := envkit.Map(e)
	fmt.Fprintf(oskit.Stdout, "Load environment: %s\n\n", m.SortedKeys())
	return envkit.Setenv(m)
}
