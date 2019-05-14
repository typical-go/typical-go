package main

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-code-generator/typigo"
	cli "gopkg.in/urfave/cli.v1"
)

func actionNewProject(c *cli.Context) (err error) {
	archetype := c.Args().Get(0)
	path := c.Args().Get(1)
	name := c.Args().Get(2)

	if archetype == "" {
		return fmt.Errorf("Archetype is missing")
	}

	if path == "" {
		return fmt.Errorf("Path is missing")
	}

	if name == "" {
		chunks := strings.Split(path, "/")
		name = chunks[len(chunks)-1]
	}

	fmt.Printf("Name=%s Archetype=%s Path=%s\n", name, archetype, path)

	return typigo.NewProject(name, archetype, path)
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
