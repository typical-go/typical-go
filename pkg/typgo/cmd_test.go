package typgo_test

import (
	"errors"
	"testing"

	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestCommand(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.Command{
				Name:            "some-name",
				Aliases:         []string{"some-alias"},
				Flags:           []cli.Flag{&cli.StringFlag{Name: "some-flag"}},
				SkipFlagParsing: true,

				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:            "some-name",
				Aliases:         []string{"some-alias"},
				Flags:           []cli.Flag{&cli.StringFlag{Name: "some-flag"}},
				SkipFlagParsing: true,
			},
			ExpectedError: "some-error",
		},
	}
	for _, tt := range testcases {
		tt.Run(t)
	}
}
