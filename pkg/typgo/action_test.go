package typgo_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
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

func TestActions(t *testing.T) {
	var seq []string
	actions := typgo.Actions{
		typgo.NewAction(func(*typgo.Context) error {
			seq = append(seq, "1")
			return nil
		}),
		typgo.NewAction(func(*typgo.Context) error {
			seq = append(seq, "2")
			return nil
		}),
	}

	require.NoError(t, actions.Execute(nil))
	require.Equal(t, []string{"1", "2"}, seq)
}

func TestContext_ExecuteBash(t *testing.T) {
	testcases := []struct {
		TestName        string
		CommandLine     string
		ExpectedErr     string
		RunExpectations []*typgo.RunExpectation
	}{
		{
			CommandLine: "some-command",
			RunExpectations: []*typgo.RunExpectation{
				{CommandLine: "some-command"},
			},
		},
		{
			CommandLine: "some-command arg1 arg2",
			RunExpectations: []*typgo.RunExpectation{
				{CommandLine: "some-command arg1 arg2"},
			},
		},
		{
			CommandLine: "",
			ExpectedErr: "command line can't be empty",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			defer typgo.PatchBash(tt.RunExpectations)(t)
			c := &typgo.Context{
				Context:  cli.NewContext(nil, &flag.FlagSet{}, nil),
				BuildSys: &typgo.BuildSys{},
			}
			err := c.ExecuteBash(tt.CommandLine)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
