package typgo

import (
	"os"
	"strings"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typlog"
	"github.com/typical-go/typical-go/pkg/typtmpl"
	"github.com/typical-go/typical-go/pkg/typvar"
	"github.com/urfave/cli/v2"
)

type (
	// BuildTool detail
	BuildTool struct {
		*Descriptor
	}

	// Context of build tool
	Context struct {
		typlog.Logger

		Cli *cli.Context
		*BuildTool
	}

	// CliFunc is command line function
	CliFunc func(*Context) error
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

// ActionFunc to return related action func
func (c *BuildTool) ActionFunc(name string, fn CliFunc) func(*cli.Context) error {
	return func(cli *cli.Context) error {
		return fn(&Context{
			Logger: typlog.Logger{
				Name: strings.ToUpper(name),
			},
			Cli:       cli,
			BuildTool: c,
		})
	}
}
