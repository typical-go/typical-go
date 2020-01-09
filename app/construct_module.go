package app

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/app/internal/tmpl"
	"github.com/typical-go/typical-go/pkg/runn"
	"github.com/typical-go/typical-go/pkg/runn/stdrun"
	"github.com/urfave/cli/v2"
)

func cmdConstructModule() *cli.Command {
	return &cli.Command{
		Name:      "module",
		Usage:     "Construct New Module",
		UsageText: "app module [Name]",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "path", Value: "pkg"},
		},
		Action: constructModule,
	}
}

func constructModule(ctx *cli.Context) (err error) {
	name := ctx.Args().First()
	if name == "" {
		return cli.ShowCommandHelp(ctx, "module")
	}
	return runn.Run(constructmodule{
		Prefix: strings.ToUpper(name),
		Name:   strings.ToLower(name),
		Path:   ctx.String("path"),
	})
}

type constructmodule struct {
	Prefix string
	Name   string
	Path   string
}

func (c constructmodule) Run() error {
	return runn.Run(
		stdrun.NewMkdir(fmt.Sprintf("%s/%s", c.Path, c.Name)),
		stdrun.NewWriteTemplate(c.path(c.Name+".go"), tmpl.Module, c),
	)
}

func (c constructmodule) path(s string) string {
	return fmt.Sprintf("%s/%s/%s", c.Path, c.Name, s)
}
