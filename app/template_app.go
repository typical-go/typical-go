package app

const appSrc = `package app

import (
	"fmt"
	"{{.Pkg}}/app/config"

	"github.com/typical-go/typical-go/pkg/typcfg"
)

// Module of application
type Module struct {}

// Action of application
func (*Module) Action() interface{} {
	return func(cfg config.Config) {
		fmt.Printf("Hello %s\n", cfg.Hello)
	}
}

// Configure the application
func (*Module) Configure() (prefix string, spec, loadFn interface{}) {
	prefix = "APP"
	spec = &config.Config{}
	loadFn = func(loader typcfg.Loader) (cfg config.Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}
`

const appSrcTest = `package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typmodule"

	"{{.Pkg}}/app"
)

func TestModule(t *testing.T) {
	a := &app.Module{}
	require.True(t, typmodule.IsActionable(a))
	require.True(t, typcfg.IsConfigurer(a))
}
`

const configSrc = "package config\n\n// Config of app\ntype Config struct {\n	Hello string `default:\"World\"`\n}"
