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
	testCmd := &typgo.TestCmd{}
	command := testCmd.Command(&typgo.BuildSys{})
	require.Equal(t, "test", command.Name)
	require.Equal(t, "Test the project", command.Usage)
	require.Equal(t, []string{"t"}, command.Aliases)
	require.NoError(t, command.Action(&cli.Context{}))
}

func TestTestCmd_Predefined(t *testing.T) {
	testCmd := &typgo.TestCmd{
		Action: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("some-error")
		}),
	}
	command := testCmd.Command(&typgo.BuildSys{})
	require.EqualError(t, command.Action(&cli.Context{}), "some-error")
}

func TestStdTest(t *testing.T) {
	stdtest := &typgo.StdTest{}
	c := &typgo.Context{
		Context: &cli.Context{Context: context.Background()},
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectLayouts: []string{"pkg3", "pkg4"}},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"go", "test", "-timeout=25s", "-coverprofile=cover.out", "./pkg3/...", "./pkg4/..."}},
	})
	defer unpatch(t)

	require.NoError(t, stdtest.Execute(c))
}

func TestStdTest_NoProjectLayout(t *testing.T) {
	stdtest := &typgo.StdTest{}
	c := &typgo.Context{
		Context: &cli.Context{Context: context.Background()},
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
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
			Descriptor: &typgo.Descriptor{ProjectLayouts: []string{"pkg3", "pkg4"}},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"go", "test", "-timeout=2m3s", "-coverprofile=some-profile", "pkg1", "pkg2"}},
	})
	defer unpatch(t)

	require.NoError(t, stdtest.Execute(c))
}
