package typbuildtool

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typenv"

	"github.com/urfave/cli/v2"
)

func (t buildtool) cmdReadme() *cli.Command {
	return &cli.Command{
		Name:   "readme",
		Usage:  "Generate readme document",
		Action: t.generateReadme,
	}
}

func (t buildtool) generateReadme(ctx *cli.Context) (err error) {
	var file *os.File
	if file, err = os.Create(typenv.Readme); err != nil {
		return
	}
	defer file.Close()
	return t.ReadmeGenerator.Generate(t.Context, file)
}
