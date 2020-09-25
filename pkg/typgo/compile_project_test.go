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

func TestCompileProject_Command(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.AppName=some-name -X github.com/typical-go/typical-go/pkg/typgo.AppVersion=0.0.1\" -o bin/some-name ./cmd/some-name"},
	})
	defer unpatch(t)

	var out strings.Builder
	typgo.Stdout = &out
	defer func() { typgo.Stdout = os.Stdout }()

	cmpl := &typgo.CompileProject{}
	s := &typgo.BuildSys{
		Descriptor: &typgo.Descriptor{AppName: "some-name", AppVersion: "0.0.1"},
	}
	command := cmpl.Command(s)
	require.Equal(t, "compile", command.Name)
	require.Equal(t, []string{"c"}, command.Aliases)
	require.Equal(t, "Compile the project", command.Usage)
	require.NoError(t, command.Action(&cli.Context{Context: context.Background()}))

	require.Equal(t, "\n$ go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.AppName=some-name -X github.com/typical-go/typical-go/pkg/typgo.AppVersion=0.0.1\" -o bin/some-name ./cmd/some-name\n", out.String())
}

func TestStdCompile_Predefined(t *testing.T) {
	cmpl := &typgo.CompileProject{
		MainPackage: "some-package",
		Output:      "some-output",
		Ldflags: execkit.BuildVars{
			"some-var": "some-value",
		},
	}
	c := &typgo.Context{
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{AppName: "some-name", AppVersion: "0.0.1"},
		},
		Context: &cli.Context{Context: context.Background()},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "go build -ldflags \"-X some-var=some-value\" -o some-output some-package"},
	})
	defer unpatch(t)

	require.NoError(t, cmpl.Execute(c))
}
