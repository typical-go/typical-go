package app

const moduleSrc = `package {{.Name}}

import (
	"github.com/typical-go/typical-go/pkg/typobj"
	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/urfave/cli/v2"
)

// Config of {{.Name}}
type Config struct {
	// TODO:
}

// Module of {{.Name}}
type Module struct {}

// Configure the module
func (m *Module) Configure() (prefix string, spec, loadFn interface{}) {
	prefix = "{{.Prefix}}"
	spec = &Config{}
	loadFn = func(loader typobj.Loader) (cfg Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}

// Provide the dependencies
func (m *Module) Provide() []interface{} {
	return []interface{}{
		// TODO: (1) put functions to be provided as dependencies
		// TODO: (2) remove this function if not required
	}
}

// Prepare the module
func (m *Module) Prepare() []interface{} {
	return []interface{}{
		// TODO: (1) put functions that run before the application start
		// TODO: (2) remove this function if not required
	}
}

// Destroy the dependencies
func (m *Module) Destroy() []interface{} {
	return []interface{}{
		// TODO: (1) functions to destroy dependencies after the application stop
		// TODO: (2) remove this function if not required
	}
}

// BuildCommands is commands to exectuce from Build-Tool
func (m *Module) BuildCommands(c *typcli.BuildCli) []*cli.Command {
	return []*cli.Command{
		// TODO: (1) add command to execute from Build-Tool
		// TODO: (2) remove this function if not required
	}
}

`

const moduleSrcTest = `package {{.Name}}

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typobj"
	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/typical-go/typical-go/pkg/typobj"
)

func TestModule(t *testing.T) {
	m := &Module{}
	require.True(t, typobj.IsProvider(m))
	require.True(t, typobj.IsDestroyer(m))
	require.True(t, typobj.IsProvider(m))
	require.True(t, typobj.IsConfigurer(m))
	require.True(t, typcli.IsBuildCommander(m))
}
`
