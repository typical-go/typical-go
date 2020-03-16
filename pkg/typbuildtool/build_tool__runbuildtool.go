package typbuildtool

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// RunBuildTool to run the build-tool
func (b *TypicalBuildTool) RunBuildTool(c *typcore.Context) (err error) {
	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = "" // NOTE: intentionally blank
	app.Description = c.Description
	app.Before = func(cliCtx *cli.Context) (err error) {
		return b.Precondition(b.createContext(c, cliCtx))
	}
	app.Version = c.Version
	app.Commands = b.Commands(c)

	return app.Run(os.Args)
}
