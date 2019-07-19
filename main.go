package main

import (
	"log"
	"os"

	"github.com/typical-go/typical-go/command"

	cli "gopkg.in/urfave/cli.v1"
)

const (
	version = "0.0.1"
)

func main() {
	app := cli.NewApp()
	app.Version = version
	app.Commands = []cli.Command{
		cli.Command{
			Name:   "new",
			Usage:  "create typical project",
			Action: newProject,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "parentPath",
					Usage: "set parent path for project path creation",
					Value: ".",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func newProject(ctx *cli.Context) error {
	projectName := ctx.Args().First()

	if projectName == "" {
		projectName = "."
	}

	return command.NewProject(ctx.String("parentPath"), projectName)

}
