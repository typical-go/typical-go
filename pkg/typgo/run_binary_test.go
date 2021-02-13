package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestRunCmd(t *testing.T) {
	runCmd := &typgo.RunBinary{}
	command := runCmd.Task().CliCommand(&typgo.Descriptor{})
	require.Equal(t, "run", command.Name)
	require.Equal(t, []string{"r"}, command.Aliases)
	require.Equal(t, "Run the project", command.Usage)
	require.True(t, command.SkipFlagParsing)

}

func TestRunCmd_Before(t *testing.T) {
	runCmd := &typgo.RunBinary{
		Before: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("before-error")
		}),
	}
	command := runCmd.Task().CliCommand(&typgo.Descriptor{})
	require.EqualError(t, command.Before(&cli.Context{}), "before-error")
}

func TestRunBinary_Execute(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "bin/some-project"},
	})(t)

	stdRun := &typgo.RunBinary{}

	require.NoError(t, stdRun.Execute(typgo.DummyContext()))
}

func TestRunBinary_Execute_Predefined(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "some-binary"},
	})(t)

	stdRun := &typgo.RunBinary{
		Binary: "some-binary",
	}
	c := typgo.DummyContext()

	require.NoError(t, stdRun.Execute(c))
}
