package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestAction(t *testing.T) {
	testcases := []struct {
		testName string
		typgo.Action
		context     *typgo.Context
		expectedErr string
	}{
		{
			Action:      typgo.NewAction(func(*typgo.Context) error { return errors.New("some-error") }),
			expectedErr: "some-error",
		},
		{
			Action: typgo.NewAction(func(*typgo.Context) error { return nil }),
		},
		{
			Action: typgo.Actions{
				typgo.NewAction(func(*typgo.Context) error { return nil }),
				typgo.NewAction(func(*typgo.Context) error { return errors.New("some-error") }),
			},
			expectedErr: "some-error",
		},
		{
			Action: typgo.Actions{
				typgo.NewAction(func(*typgo.Context) error { return errors.New("some-error") }),
				typgo.NewAction(func(*typgo.Context) error { return nil }),
			},
			expectedErr: "some-error",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			err := tt.Execute(tt.context)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

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
