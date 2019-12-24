package app

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

const (
	// Version of Typical-Go
	Version = "0.9.20"
)

// Module of app
type Module struct{}

// AppCommands is commands collection to execute application
func (m Module) AppCommands(c *typcore.Context) []*cli.Command {
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
// 	loadFn = func(loader typcore.ConfigLoader) (cfg Config, err error) {
// 		err = loader.Load(prefix, &cfg)
// 		return
// 	}
// 	return
// }

// type Config struct {
// 	Hello string `default:"world"`
// }
