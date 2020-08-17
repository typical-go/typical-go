package typgo_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestCleanProject(t *testing.T) {
	typgo.TypicalTmp = "some-tmp"
	cleanPrj := &typgo.CleanProject{}

	var out strings.Builder
	typgo.Stdout = &out
	defer func() {
		typgo.Stdout = os.Stdout
	}()

	command := cleanPrj.Command(&typgo.BuildSys{})
	require.Equal(t, "clean", command.Name)
	require.Equal(t, "Clean the project", command.Usage)
	require.NoError(t, command.Action(&cli.Context{}))

	require.Equal(t, "Removing some-tmp\n", out.String())

}

func TestCleanProject_Execute(t *testing.T) {
	var out strings.Builder
	typgo.Stdout = &out
	defer func() {
		typgo.Stdout = os.Stdout
	}()

	cleanPrj := &typgo.CleanProject{
		Paths: []string{"path1", "path2"},
	}
	require.NoError(t, cleanPrj.Execute(&typgo.Context{}))
	require.Equal(t, "Removing path1\nRemoving path2\n", out.String())
}
