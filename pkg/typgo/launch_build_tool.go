package typgo

import (
	"os"

	"github.com/urfave/cli/v2"
)

func launchBuildTool(d *Descriptor) error {
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "Build-Tool"
	app.Description = d.Description
	app.Version = d.Version
	app.Before = beforeBuildTool(d)

	app.Commands = createBuildToolCmds(d)

	return app.Run(os.Args)
}
