package typicalgo

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/common/stdrun"
	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/typicalgo/internal/tmpl"
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
	return common.Run(
		wrapper(c.String("path"), pkg),
	)
}

func wrapper(path, pkg string) common.Runner {
	return stdrun.NewWriteTemplate(
		path+"/typicalw",
		tmpl.Typicalw,
		tmpl.TypicalwData{
			DescriptorPackage: fmt.Sprintf("%s/typical", pkg),
			DescriptorFile:    "typical/descriptor.go",
			ChecksumFile:      ".typical-tmp/checksum",
			LayoutTemp:        typcore.DefaultLayout.Temp,
		},
	).WithPermission(0700)
}
