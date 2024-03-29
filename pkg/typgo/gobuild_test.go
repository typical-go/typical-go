package typgo_test

import (
	"flag"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestGoBuild_Command(t *testing.T) {

	cmpl := &typgo.GoBuild{}

	command := cmpl.Task().CliCommand(&typgo.Descriptor{})
	require.Equal(t, "build", command.Name)
	require.Equal(t, []string{"b"}, command.Aliases)
	require.Equal(t, "build the project", command.Usage)

	var out strings.Builder
	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
		Logger:  typgo.Logger{Stdout: &out},
		Descriptor: &typgo.Descriptor{
			ProjectName:    "some-project",
			ProjectVersion: "0.0.1",
		},
	}
	defer c.PatchBash([]*typgo.MockCommand{
		{CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=0.0.1\" -o bin/some-project ./cmd/some-project"},
	})(t)

	require.NoError(t, cmpl.Execute(c))
	require.Equal(t, "> go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=0.0.1\" -o bin/some-project ./cmd/some-project\n", out.String())
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
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
	}
	defer c.PatchBash([]*typgo.MockCommand{
		{CommandLine: "go build -ldflags \"-X some-var=some-value\" -o some-output some-package"},
	})(t)

	require.NoError(t, cmpl.Execute(c))
}
