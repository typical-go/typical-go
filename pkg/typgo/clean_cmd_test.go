package typgo_test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestCleanCmd(t *testing.T) {
	cleanCmd := &typgo.CleanCmd{}
	command := cleanCmd.Command(&typgo.BuildSys{})
	require.Equal(t, "clean", command.Name)
	require.Equal(t, "Clean the project", command.Usage)
	require.NoError(t, command.Action(&cli.Context{}))
}

func TestCleanCmd_Predefined(t *testing.T) {
	cleanCmd := &typgo.CleanCmd{
		Action: typgo.NewAction(func(*typgo.Context) error {
			return errors.New("some-error")
		}),
	}
	command := cleanCmd.Command(&typgo.BuildSys{})
	require.EqualError(t, command.Action(&cli.Context{}), "some-error")
}

func TestStdClean(t *testing.T) {
	typgo.TypicalTmp = "some-tmp"
	stdClean := &typgo.StdClean{}

	var out strings.Builder
	typgo.Stdout = &out
	defer func() {
		typgo.Stdout = os.Stdout
	}()

	require.NoError(t, stdClean.Execute(&typgo.Context{}))
	require.Equal(t, "Removing some-tmp\n", out.String())

}

func TestStdClean_Predefined(t *testing.T) {
	var out strings.Builder
	typgo.Stdout = &out
	defer func() {
		typgo.Stdout = os.Stdout
	}()

	stdClean := &typgo.StdClean{
		Paths: []string{"path1", "path2"},
	}
	require.NoError(t, stdClean.Execute(&typgo.Context{}))
	require.Equal(t, "Removing path1\nRemoving path2\n", out.String())
}
