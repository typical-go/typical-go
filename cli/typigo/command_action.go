package main

import (
	"fmt"
	"log"

	"github.com/typical-go/typical-code-generator/typigo"
	cli "gopkg.in/urfave/cli.v1"
)

func actionNewProject(c *cli.Context) (err error) {

	path := c.Args().Get(0)

	if path == "" {
		return fmt.Errorf("Path is missing")
	}

	log.Printf("New Project at '%s'\n", path)
	return typigo.NewProject(path)
}

func actionContext(c *cli.Context) (err error) {
	field := c.Args().Get(0)
	value := c.Args().Get(1)

	if field == "" {
		// TODO: print all context
		return
	}

	if value == "" {
		// TODO: print specific field
		return
	}

	// TODO: set field
	return
}
