package app

import (
	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/urfave/cli/v2"
)

const (
	// Version of Typical-Go
	Version = "0.9.5"
)

// Module of application
func Module() interface{} {
	return &module{}
}

type module struct{}

func (module) AppCommands(c typcli.Cli) []*cli.Command {
	return []*cli.Command{
		cmdConstructProject(),
		cmdConstructModule(),
		cmdCreateWrapper(),
	}
}
