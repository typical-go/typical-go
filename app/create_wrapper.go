package app

import (
	"github.com/typical-go/typical-go/pkg/utility/runn"
	"github.com/typical-go/typical-go/pkg/utility/runner"
	"github.com/urfave/cli"
)

func cmdCreateWrapper() cli.Command {
	return cli.Command{
		Name:  "wrapper",
		Usage: "Create the wrapper",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "path", Value: "."},
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
	return runner.WriteString{
		Target:     path + "/typicalw",
		Permission: 0700,
		Content:    typicalw,
	}
}
