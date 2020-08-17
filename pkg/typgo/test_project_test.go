package typgo_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestTestProject(t *testing.T) {
	testPrj := &typgo.TestProject{}

	c := &cli.Context{Context: context.Background()}
	sys := &typgo.BuildSys{
		Descriptor: &typgo.Descriptor{ProjectLayouts: []string{"pkg3", "pkg4"}},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"go", "test", "-timeout=25s", "-coverprofile=cover.out", "./pkg3/...", "./pkg4/..."}},
	})
	defer unpatch(t)

	command := testPrj.Command(sys)
	require.Equal(t, "test", command.Name)
	require.Equal(t, "Test the project", command.Usage)
	require.Equal(t, []string{"t"}, command.Aliases)
	require.NoError(t, command.Action(c))
}

func TestTestProject_NoProjectLayout(t *testing.T) {
	testProj := &typgo.TestProject{}
	c := &typgo.Context{
		Context: &cli.Context{Context: context.Background()},
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	require.NoError(t, testProj.Execute(c))
}

func TestTestProject_Predefined(t *testing.T) {
	testProj := &typgo.TestProject{
		Timeout:      123 * time.Second,
		CoverProfile: "some-profile",
		Packages:     []string{"pkg1", "pkg2"},
	}

	c := &typgo.Context{
		Context: &cli.Context{Context: context.Background()},
		BuildSys: &typgo.BuildSys{
			Descriptor: &typgo.Descriptor{ProjectLayouts: []string{"pkg3", "pkg4"}},
		},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"go", "test", "-timeout=2m3s", "-coverprofile=some-profile", "pkg1", "pkg2"}},
	})
	defer unpatch(t)

	require.NoError(t, testProj.Execute(c))
}
