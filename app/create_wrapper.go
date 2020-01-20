package app

import (
	"fmt"

	"github.com/typical-go/typical-go/app/internal/tmpl"
	"github.com/typical-go/typical-go/pkg/runn"
	"github.com/typical-go/typical-go/pkg/runn/stdrun"
	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/urfave/cli/v2"
)

func cmdCreateWrapper() *cli.Command {
	return &cli.Command{
		Name:  "wrapper",
		Usage: "Create the wrapper",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "path", Value: "."},
		},
		Action: createWrapper,
	}
}

func createWrapper(c *cli.Context) error {
	pkg := c.Args().First()
	if pkg == "" {
		return cli.ShowCommandHelp(c, "wrapper")
	}
	return runn.Run(
		wrapper(c.String("path"), pkg),
	)
}

func wrapper(path, pkg string) runn.Runner {
	return stdrun.NewWriteTemplate(
		path+"/typicalw",
		tmpl.Typicalw,
		tmpl.TypicalwData{
			DescriptorPackage: fmt.Sprintf("%s/typical", pkg),
			DescriptorFile:    typenv.DescriptorFile,
			ChecksumFile:      typenv.ChecksumFile,
			LayoutTemp:        typenv.Layout.Temp,
		},
	).WithPermission(0700)
}
