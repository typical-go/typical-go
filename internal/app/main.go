package app

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

// Main function the typical-go
func Main() (err error) {
	return App().Run(os.Args)
}

// App application
func App() *cli.App {
	app := cli.NewApp()
	app.Name = typgo.AppName
	app.Version = typgo.AppVersion
	app.Usage = ""       // NOTE: intentionally blank
	app.Description = "" // NOTE: intentionally blank
	app.Commands = []*cli.Command{
		cmdRun(),
		cmdSetup(),
	}
	return app
}
