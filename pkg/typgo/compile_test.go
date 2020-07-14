package typgo_test

import (
	"errors"
	"testing"

	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestCompileCmd(t *testing.T) {
	testcases := []typgo.CmdTestCase{
		{
			Cmd: &typgo.CompileCmd{
				Action: typgo.NewAction(func(*typgo.Context) error {
					return errors.New("some-error")
				}),
			},
			Expected: typgo.Command{
				Name:    "compile",
				Aliases: []string{"c"},
				Usage:   "Compile the project",
			},
			ExpectedError: "some-error",
		},
	}

	for _, tt := range testcases {
		tt.Run(t)
	}
}
