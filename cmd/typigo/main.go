package main

import (
	"log"
	"os"

	cli "gopkg.in/urfave/cli.v1"
)

const (
	Version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Commands = []cli.Command{
		cli.Command{Name: "new", Action: actionNewProject},
		cli.Command{Name: "context", ShortName: "ctx", Action: actionContext},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
