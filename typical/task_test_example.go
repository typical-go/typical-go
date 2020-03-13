package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/urfave/cli/v2"
)

func taskTestExample(ctx *typbuildtool.Context) *cli.Command {
	return &cli.Command{
		Name:    "test-example",
		Aliases: []string{"e"},
		Usage:   "Test all example",
		Action: func(cliCtx *cli.Context) (err error) {
			gotest := buildkit.NewGoTest("./examples/...")
			cmd := gotest.Command(cliCtx.Context)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			return cmd.Run()
		},
	}
}
