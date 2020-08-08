package app

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/urfave/cli/v2"
)

// Start the typical-go
func Start() (err error) {
	app := cli.NewApp()
	app.Name = typapp.Name
	app.Version = typapp.Version
	app.Usage = ""       // NOTE: intentionally blank
	app.Description = "" // NOTE: intentionally blank
	app.Commands = []*cli.Command{
		cmdRun(),
		cmdSetup(),
	}
	return app.Run(os.Args)
}
