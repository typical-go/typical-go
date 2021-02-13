package typgo_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/oskit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestGoBuild_Command(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=0.0.1\" -o bin/some-project ./cmd/some-project"},
	})(t)

	var out strings.Builder
	defer oskit.PatchStdout(&out)()

	cmpl := &typgo.GoBuild{}

	command := cmpl.Task().CliCommand(&typgo.Descriptor{})
	require.Equal(t, "build", command.Name)
	require.Equal(t, []string{"b"}, command.Aliases)
	require.Equal(t, "build the project", command.Usage)

	require.NoError(t, cmpl.Execute(typgo.DummyContext()))
	require.Equal(t, "some-project:dummy> $ go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-project -X github.com/typical-go/typical-go/pkg/typgo.ProjectVersion=0.0.1\" -o bin/some-project ./cmd/some-project\n", out.String())
}

func TestGoBuild_Predefined(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "go build -ldflags \"-X some-var=some-value\" -o some-output some-package"},
	})(t)

	cmpl := &typgo.GoBuild{
		MainPackage: "some-package",
		Output:      "some-output",
		Ldflags: typgo.BuildVars{
			"some-var": "some-value",
		},
	}
	c := typgo.DummyContext()

	require.NoError(t, cmpl.Execute(c))
}
