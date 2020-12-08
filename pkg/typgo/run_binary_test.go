package typgo_test

import (
	"errors"
	"flag"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestRunCmd(t *testing.T) {
	runCmd := &typgo.RunBinary{}
	command := runCmd.Task(&typgo.BuildSys{Descriptor: &typgo.Descriptor{}})
	require.Equal(t, "run", command.Name)
	require.Equal(t, []string{"r"}, command.Aliases)
	require.Equal(t, "Run the project", command.Usage)
	require.True(t, command.SkipFlagParsing)

	c := cli.NewContext(nil, &flag.FlagSet{}, nil)
	require.NoError(t, command.Before(c))
}

func TestRunCmd_Before(t *testing.T) {
	runCmd := &typgo.RunBinary{
		Before: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("before-error")
		}),
	}
	command := runCmd.Task(&typgo.BuildSys{})
	require.EqualError(t, command.Before(&cli.Context{}), "before-error")
}

func TestRunBinary_Execute(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "bin/some-name"},
	})(t)

	stdRun := &typgo.RunBinary{}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectName: "some-name"},
		},
	}

	require.NoError(t, stdRun.Execute(c))
}

func TestRunBinary_Execute_Predefined(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "some-binary"},
	})(t)

	stdRun := &typgo.RunBinary{
		Binary: "some-binary",
	}
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectName: "some-name"},
		},
	}

	require.NoError(t, stdRun.Execute(c))
}
