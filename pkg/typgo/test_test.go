package typgo_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
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

func TestStdTest(t *testing.T) {

	stdtest := &typgo.StdTest{}
	c := &typgo.Context{
		Context: &cli.Context{Context: context.Background()},
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{Layouts: []string{"pkg3", "pkg4"}},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"go", "test", "-timeout=25s", "-coverprofile=cover.out", "./pkg3/...", "./pkg4/..."}},
	})
	defer unpatch(t)

	require.NoError(t, stdtest.Execute(c))
}

func TestStdTest_Predefined(t *testing.T) {
	stdtest := &typgo.StdTest{
		Timeout:      123 * time.Second,
		CoverProfile: "some-profile",
		Packages:     []string{"pkg1", "pkg2"},
	}

	c := &typgo.Context{
		Context: &cli.Context{Context: context.Background()},
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{Layouts: []string{"pkg3", "pkg4"}},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"go", "test", "-timeout=2m3s", "-coverprofile=some-profile", "pkg1", "pkg2"}},
	})
	defer unpatch(t)

	require.NoError(t, stdtest.Execute(c))
}
