package app

import (
	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/urfave/cli/v2"
)

const (
	// Version of Typical-Go
	Version = "0.9.8"
)

// Module of app
type Module struct{}

// Commands of module
func (Module) Commands(c *typcli.AppCli) []*cli.Command {
	return []*cli.Command{
		cmdConstructProject(),
		cmdConstructModule(),
		cmdCreateWrapper(),
	}
}

// -- Uncomment to test action --
// func (Module) Action() interface{} {
// 	return func(cfg Config) {
// 		fmt.Printf("Hello %s\n", cfg.Hello)
// 	}
// }

// func (Module) Configure() (prefix string, spec, loadFn interface{}) {
// 	prefix = "APP"
// 	spec = &Config{}
// 	loadFn = func(loader typcfg.Loader) (cfg Config, err error) {
// 		err = loader.Load(prefix, &cfg)
// 		return
// 	}
// 	return
// }

// type Config struct {
// 	Hello string `default:"world"`
// }
