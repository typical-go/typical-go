package app

import (
	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/urfave/cli/v2"
)

const (
	// Version of Typical-Go
	Version = "0.9.7"
)

// Module of application
func Module() interface{} {
	return &module{}
}

type module struct{}

func (module) Commands(c *typcli.AppCli) []*cli.Command {
	return []*cli.Command{
		cmdConstructProject(),
		cmdConstructModule(),
		cmdCreateWrapper(),
	}
}
