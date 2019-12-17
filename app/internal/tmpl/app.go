package tmpl

// App template
const App = `package app

import (
	"fmt"
	"{{.Pkg}}/app/config"
	"github.com/typical-go/typical-go/pkg/typcore"
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
	loadFn = func(loader typcore.ConfigLoader) (cfg config.Config, err error) {
		err = loader.Load(prefix, &cfg)
		return
	}
	return
}
`
