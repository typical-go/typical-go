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

func TestRunCmd_1(t *testing.T) {
	runCmd := &typgo.RunCmd{
		Before: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("before-error")
		}),
		Action: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("action-error")
		}),
	}
	command := runCmd.Command(&typgo.BuildSys{})
	require.EqualError(t, command.Action(&cli.Context{}), "action-error")
	require.EqualError(t, command.Before(&cli.Context{}), "before-error")
}

func TestRunProject_Execute(t *testing.T) {
	stdRun := &typgo.RunProject{}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{AppName: "some-name"},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "bin/some-name"},
	})
	defer unpatch(t)

	require.NoError(t, stdRun.Execute(c))
}

func TestRunProject_Execute_Predefined(t *testing.T) {
	stdRun := &typgo.RunProject{
		Binary: "some-binary",
	}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{AppName: "some-name"},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "some-binary"},
	})
	defer unpatch(t)

	require.NoError(t, stdRun.Execute(c))
}
