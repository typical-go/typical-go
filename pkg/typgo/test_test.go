package typgo_test

import (
	"errors"
	"testing"

	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestTestCmd(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.TestCmd{
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "Test the project",
			},
			ExpectedError: "some-error",
		},
	}
	for _, tt := range testcases {
		tt.Run(t)
	}

}
