package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/exor"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

func taskTestExample(ctx *typbuildtool.Context) *cli.Command {
	return &cli.Command{
		Name:    "test-example",
		Aliases: []string{"e"},
		Usage:   "Test all example",
		Action: func(cliCtx *cli.Context) (err error) {
			gotest := exor.NewGoTest("./examples/...").
				WithStdout(os.Stdout).
				WithStderr(os.Stderr)

			return gotest.Execute(cliCtx.Context)
		},
	}
}
