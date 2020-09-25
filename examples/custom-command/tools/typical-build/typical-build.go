package main

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/typgo"
)

var descriptor = typgo.Descriptor{
	AppName:    "custom-command",
	AppVersion: "1.0.0",

	Cmds: []typgo.Cmd{
		// compile
		&typgo.CompileProject{},
		// clean
		&typgo.CleanProject{},
		// run
		&typgo.RunCmd{
			Before: typgo.BuildCmdRuns{"compile"},
			Action: &typgo.RunProject{},
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
