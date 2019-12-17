package app

import (
	"fmt"
	"strings"

	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/typical-go/typical-go/pkg/utility/runner"
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
	return runn.Execute(constructmodule{
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
	return runn.Execute(
		runner.NewMkdir(fmt.Sprintf("%s/%s", c.Path, c.Name)),
		runner.NewWriteTemplate(c.path(c.Name+".go"), moduleSrc, c),
		runner.NewWriteTemplate(c.path(c.Name+"_test.go"), moduleSrcTest, c),
	)
}

func (c constructmodule) path(s string) string {
	return fmt.Sprintf("%s/%s/%s", c.Path, c.Name, s)
}
