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

func (b *BuildTool) cli(tc *typcore.Context) *cli.App {
	c := b.context(tc)

	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = "Build-Tool"
	app.Description = c.Description

	app.Before = b.before(c)
	app.Version = c.Version
	app.Commands = b.Commands(c)

	return app
}

func (b *BuildTool) context(tc *typcore.Context) *Context {
	return &Context{
		Context:   tc,
		BuildTool: b,
	}
}