package app

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/typical-go/typical-go/pkg/typenv"
	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/typical-go/typical-go/pkg/utility/runner"
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
	return runn.Execute(
		wrapperRunner(ctx.String("path")),
	)
}

func wrapperRunner(path string) runn.Runner {
	var name string
	if path == "." {
		dir, _ := os.Getwd()
		name = filepath.Base(dir)
	} else {
		name = filepath.Base(path)
	}

	return runner.WriteTemplate{
		Target:     path + "/typicalw",
		Permission: 0700,
		Template:   typicalw,
		Data: struct {
			ContextFile        string
			ChecksumFile       string
			LayoutMetadata     string
			PrebuilderBin      string
			PrebuilderMainPath string
			BuildtoolBin       string
		}{
			ContextFile:        typenv.ContextFile,
			ChecksumFile:       typenv.ChecksumFile,
			LayoutMetadata:     typenv.Layout.Metadata,
			PrebuilderBin:      fmt.Sprintf("%s/%s-%s", typenv.Layout.Bin, name, typenv.Prebuilder),
			PrebuilderMainPath: fmt.Sprintf("%s/%s-%s", typenv.Layout.Cmd, name, typenv.Prebuilder),
			BuildtoolBin:       fmt.Sprintf("%s/%s-%s", typenv.Layout.Bin, name, typenv.BuildTool),
		},
	}
}
