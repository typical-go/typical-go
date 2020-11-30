package main

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	ProjectName:    "custom-command",
	ProjectVersion: "1.0.0",

	Cmds: []typgo.CliCommander{
		// compile
		&typgo.GoBuild{},
		// run
		&typgo.RunProject{
			Before: typgo.BuildCmdRuns{"compile"},
		},
		// ping
		&typgo.Command{
			Name: "ping",
			Action: typgo.NewAction(
				func(c *typgo.Context) error {
					fmt.Println("pong")
					return nil
				},
			),
		},
	},
}

func main() {
	typgo.Start(&descriptor)
}
