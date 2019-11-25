package app

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/urfave/cli"
)

const (
	// Version of Typical-Go
	Version = "0.9.1"
)

// Module of application
func Module() interface{} {
	return &module{}
}

type module struct{}

func (module) AppCommands(c *typcli.ContextCli) []cli.Command {
	return []cli.Command{
		{Name: "init", Usage: "Iniate new project", Action: initiateProject},
		{Name: "migrate-me", Usage: "Migrate current project to using framework", Action: underConstruction},
		{Name: "upgrade", Usage: "upgrade the typical-go", Action: underConstruction},
		{Name: "update", Usage: "Update current project to use latest framework", Action: underConstruction},
	}
}

func underConstruction() {
	fmt.Println("Under Construction")
}
