package typbuildtool

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcfg"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// RunBuildTool to run the build-tool
func (b *BuildTool) RunBuildTool(tc *typcore.Context) (err error) {
	c := &Context{
		Context:   tc,
		BuildTool: b,
	}

	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = "Build-Tool"
	app.Description = c.Description

	app.Before = func(cliCtx *cli.Context) (err error) {
		if err = b.Precondition(c.BuildContext(cliCtx)); err != nil {
			return
		}

		typcfg.Load(b.configFile)

		return
	}
	app.Version = c.Version
	app.Commands = b.Commands(c)

	return app.Run(os.Args)
}
