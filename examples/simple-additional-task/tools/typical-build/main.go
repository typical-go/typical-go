package main

import (
	"fmt"
	"log"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

var (
	// Descriptor of sample
	descriptor = typgo.Descriptor{
		Name:    "simple-additional-task",
		Version: "1.0.0",
		Layouts: []string{"internal"},

		Compile: &typgo.StdCompile{},
		Run:     &typgo.StdRun{},
		Clean:   &typgo.StdClean{},

		Utility: typgo.Utilities{
			typgo.CreateUtility(&cli.Command{
				Name: "ping",
				Action: func(c *cli.Context) error {
					fmt.Println("pong")
					return nil
				},
			}),
			typgo.NewUtility(func(c *typgo.BuildCli) ([]*cli.Command, error) {
				return []*cli.Command{
					{
						Name:   "desc",
						Usage:  "Print descriptor",
						Action: c.ActionFn("PRINT_DESC", printDesc),
					},
				}, nil
			}),
		},
	}
)

func printDesc(c *typgo.Context) error {
	fmt.Printf("name=%s\n", c.Descriptor.Name)
	fmt.Printf("version=%s\n", c.Descriptor.Version)
	return nil
}

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err)
	}
}
