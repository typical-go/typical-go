package typgo_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestContext_ExecuteBash(t *testing.T) {
	testcases := []struct {
		TestName    string
		CommandLine string
		ExpectedErr string
		MockBashs   []*typgo.MockBash
	}{
		{
			CommandLine: "some-command",
			MockBashs: []*typgo.MockBash{
				{CommandLine: "some-command"},
			},
		},
		{
			CommandLine: "some-command arg1 arg2",
			MockBashs: []*typgo.MockBash{
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
			c := &typgo.Context{}
			defer c.PatchBash(tt.MockBashs)(t)
			err := c.ExecuteBash(tt.CommandLine)
			if tt.ExpectedErr != "" {
				require.EqualError(t, err, tt.ExpectedErr)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestContext_PrintInfo(t *testing.T) {
	var out strings.Builder
	c := &typgo.Context{
		Context: &cli.Context{
			Command: &cli.Command{
				Name: "some-command",
			},
		},
		Descriptor: &typgo.Descriptor{
			ProjectName: "some-project",
		},
		Stdout: &out,
	}
	c.Info("some-info")
	c.Infof("some-info %d\n", 9999)

	require.Equal(t, "some-project:some-command> some-info\nsome-project:some-command> some-info 9999\n", out.String())
}

func TestContext_PrintInfo_NoPanic(t *testing.T) {
	c := &typgo.Context{}
	c.Info("some-info")
	c.Infof("some-info %d\n", 9999)
}

func TestContext_PatchContext(t *testing.T) {
	c := &typgo.Context{}
	defer c.PatchBash([]*typgo.MockBash{
		{CommandLine: "name1 arg1", ReturnError: errors.New("some-error")},
	})(t)

	bash := &typgo.Bash{Name: "name1", Args: []string{"arg1"}}
	require.EqualError(t, c.Execute(bash), "some-error")
}
