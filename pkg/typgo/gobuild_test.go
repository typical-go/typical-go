package typgo_test

import (
	"context"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestGoBuild_Command(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-name -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=0.0.1\" -o bin/some-name ./cmd/some-name"},
	})
	defer unpatch(t)

	var out strings.Builder
	typgo.Stdout = &out
	defer func() { typgo.Stdout = os.Stdout }()

	cmpl := &typgo.GoBuild{}
	s := &typgo.BuildSys{
		Descriptor: &typgo.Descriptor{ProjectName: "some-name", ProjectVersion: "0.0.1"},
	}
	command := cmpl.Task(s)
	require.Equal(t, "build", command.Name)
	require.Equal(t, []string{"b"}, command.Aliases)
	require.Equal(t, "build the project", command.Usage)
	require.NoError(t, command.Action(&cli.Context{Context: context.Background()}))

	require.Equal(t, "\n$ go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-name -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=0.0.1\" -o bin/some-name ./cmd/some-name\n", out.String())
}

func TestGoBuild_Predefined(t *testing.T) {
	cmpl := &typgo.GoBuild{
		MainPackage: "some-package",
		Output:      "some-output",
		Ldflags: typgo.BuildVars{
			"some-var": "some-value",
		},
	}
	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectName: "some-name", ProjectVersion: "0.0.1"},
		},
		Context: &cli.Context{Context: context.Background()},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "go build -ldflags \"-X some-var=some-value\" -o some-output some-package"},
	})
	defer unpatch(t)

	require.NoError(t, cmpl.Execute(c))
}
