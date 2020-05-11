package typgo

import (
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

func launchBuildTool(d *Descriptor) error {
	app := cli.NewApp()
	app.Name = d.Name
	app.Usage = "Build-Tool"
	app.Description = d.Description
	app.Version = d.Version
	app.Before = beforeBuildTool(d)

	buildTool := &BuildTool{Descriptor: d}
	app.Commands = buildTool.Commands()

	return app.Run(os.Args)
}

func beforeBuildTool(d *Descriptor) cli.BeforeFunc {
	return func(cli *cli.Context) (err error) {
		os.Remove(typvar.PrecondFile)
		ctx := cli.Context
		c := createPrecondContext(ctx, d)

		if err = d.Precondition(c); err != nil {
			return
		}

		if len(c.Lines) > 0 {
			if err = typtmpl.WriteFile(typvar.PrecondFile, 0777, c); err != nil {
				return
			}
			if err = buildkit.GoImports(ctx, typvar.PrecondFile); err != nil {
				return
			}
		} else {
			c.Info("No precondition")
			os.Remove(typvar.PrecondFile)
		}
		return
	}
}
