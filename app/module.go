package app

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

const (
	// Version of Typical-Go
	Version = "0.9.23"
)

// Module of Typical-Go
func Module() interface{} {
	return &module{}
}

type module struct{}

func (m module) AppCommands(c *typcore.Context) []*cli.Command {
	return []*cli.Command{
		cmdConstructProject(),
		cmdConstructModule(),
		cmdCreateWrapper(),
	}
}
