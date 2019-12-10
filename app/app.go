package app

import (
	"github.com/typical-go/typical-go/pkg/typobj"
	"github.com/urfave/cli/v2"
)

const (
	// Version of Typical-Go
	Version = "0.9.11"
)

// Module of app
type Module struct{}

// AppCommands is commands collection to execute application
func (Module) AppCommands(c typobj.Cli) []*cli.Command {
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
// 	loadFn = func(loader typobj.Loader) (cfg Config, err error) {
// 		err = loader.Load(prefix, &cfg)
// 		return
// 	}
// 	return
// }

// type Config struct {
// 	Hello string `default:"world"`
// }
