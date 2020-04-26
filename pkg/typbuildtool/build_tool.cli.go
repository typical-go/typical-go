package typbuildtool

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// RunBuildTool to run the build-tool
func (b *BuildTool) RunBuildTool(tc *typcore.Context) (err error) {
	return b.cli(tc).Run(os.Args)
}

func (b *BuildTool) cli(core *typcore.Context) *cli.App {

	app := cli.NewApp()
	app.Name = core.Name
	app.Usage = "Build-Tool"
	app.Description = core.Description

	c := b.context(core)
	app.Before = c.ActionFunc(b.Precondition)
	app.Version = c.Core.Version
	app.Commands = b.Commands(c)

	return app
}

func (b *BuildTool) context(core *typcore.Context) *Context {
	return &Context{
		Core:      core,
		BuildTool: b,
	}
}
