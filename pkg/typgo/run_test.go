package typgo_test

import (
	"errors"
	"testing"

	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestRunCompile(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.RunCmd{
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:            "run",
				Aliases:         []string{"r"},
				Usage:           "Run the project in local environment",
				SkipFlagParsing: true,
			},
			ExpectedError: "some-error",
		},
	}
	for _, tt := range testcases {
		tt.Run(t)
	}
}
