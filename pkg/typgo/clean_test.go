package typgo_test

import (
	"errors"
	"testing"

	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCleanCmd(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.CleanCmd{
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:  "clean",
				Usage: "Clean the project",
			},
			ExpectedError: "some-error",
		},
	}
	for _, tt := range testcases {
		tt.Run(t)
	}

}
