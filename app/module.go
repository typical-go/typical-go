package app

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

const (
	// Version of Typical-Go
	Version = "0.9.26"
)

// New of Typical-Go
func New() *Module {
	return &Module{}
}

// Module of Typical-Go
type Module struct{}

// AppCommands return command
func (m Module) AppCommands(a *typcore.AppContext) []*cli.Command {
	return []*cli.Command{
		cmdConstructProject(),
		cmdConstructModule(),
		cmdCreateWrapper(),
	}
}
