package app

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/urfave/cli"
)

// Module of application
func Module() typctx.AppModule {
	return &module{}
}

type module struct{}

func (module) Run() interface{} {
	return underConstruction
}

func (module) AppCommands(c *typcli.ContextCli) []cli.Command {
	return []cli.Command{
		{Name: "init", Usage: "Iniate new project", Action: underConstruction},
		{Name: "migrate-me", Usage: "Migrate current project to using framework", Action: underConstruction},
		{Name: "upgrade", Usage: "upgrade the typical-go", Action: underConstruction},
		{Name: "update", Usage: "Update current project to use latest framework", Action: underConstruction},
	}
}

func underConstruction() {
	fmt.Println("Under Construction")
}
