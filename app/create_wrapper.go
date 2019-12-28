package app

import (
	"fmt"
	"os"
	"path/filepath"

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

func createWrapper(ctx *cli.Context) error {
	return runn.Run(
		wrapper(ctx.String("path")),
	)
}

func wrapper(path string) runn.Runner {
	var name string
	if path == "." {
		dir, _ := os.Getwd()
		name = filepath.Base(dir)
	} else {
		name = filepath.Base(path)
	}
	data := struct {
		DescriptorFile    string
		ChecksumFile      string
		LayoutTemp        string
		BuildtoolMainPath string
		BuildtoolBin      string
	}{
		DescriptorFile:    typenv.DescriptorFile,
		ChecksumFile:      typenv.ChecksumFile,
		LayoutTemp:        typenv.Layout.Temp,
		BuildtoolMainPath: fmt.Sprintf("%s/%s-%s", typenv.Layout.Cmd, name, typenv.BuildTool),
		BuildtoolBin:      fmt.Sprintf("%s/%s-%s", typenv.Layout.Bin, name, typenv.BuildTool),
	}
	return stdrun.NewWriteTemplate(path+"/typicalw", tmpl.Typicalw, data).WithPermission(0700)
}
