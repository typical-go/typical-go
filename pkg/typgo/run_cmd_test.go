package typgo_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestRunCmd(t *testing.T) {
	runCmd := &typgo.RunCmd{}
	command := runCmd.Command(&typgo.BuildSys{})
	require.Equal(t, "run", command.Name)
	require.Equal(t, []string{"r"}, command.Aliases)
	require.Equal(t, "Run the project", command.Usage)
	require.True(t, command.SkipFlagParsing)
	require.NoError(t, command.Action(&cli.Context{}), "some-error")
	require.NoError(t, command.Before(&cli.Context{}))
}

func TestRunCommand_(t *testing.T) {
	runCmd := &typgo.RunCmd{
		Name:    "some-name",
		Usage:   "some-usage",
		Aliases: []string{"x"},
		Before: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("before-error")
		}),
		Action: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("action-error")
		}),
	}
	command := runCmd.Command(&typgo.BuildSys{})
	require.Equal(t, "some-name", command.Name)
	require.Equal(t, "some-usage", command.Usage)
	require.Equal(t, []string{"x"}, command.Aliases)
	require.EqualError(t, command.Action(&cli.Context{}), "action-error")
	require.EqualError(t, command.Before(&cli.Context{}), "before-error")
}

func TestStdRun_Execute(t *testing.T) {
	stdRun := &typgo.StdRun{}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectName: "some-name"},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"bin/some-name"}},
	})
	defer unpatch(t)

	require.NoError(t, stdRun.Execute(c))
}

func TestStdRun_Execute_Predefined(t *testing.T) {
	stdRun := &typgo.StdRun{
		Binary: "some-binary",
	}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectName: "some-name"},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"some-binary"}},
	})
	defer unpatch(t)

	require.NoError(t, stdRun.Execute(c))
}
