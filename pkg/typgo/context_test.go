package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

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
			c, _ := typgo.DummyContext()
			err := c.ExecuteBash(tt.CommandLine)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}

}
