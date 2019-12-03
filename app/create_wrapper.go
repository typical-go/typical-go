package app

import (
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
	return runner.WriteTemplate{
		Target:   path + "/typicalw",
		Template: typicalw,
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
			PrebuilderBin:      typenv.PrebuilderBin,
			PrebuilderMainPath: typenv.PrebuilderMainPath,
			BuildtoolBin:       typenv.BuildToolBin,
		},
	}
}
