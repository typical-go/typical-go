package typgo_test

import (
	"context"
	"path/filepath"
	"testing"

	"github.com/typical-go/typical-go/pkg/filekit"

	"bou.ke/monkey"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestTestProject(t *testing.T) {
	defer monkey.Patch(filepath.Walk,
		func(root string, walkFn filepath.WalkFunc) error {
			walkFn("pkg1", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("pkg2", &filekit.FileInfo{IsDirField: true}, nil)
			return nil
		},
	).Unpatch()
	defer execkit.Patch(nil)(t)

	c := &cli.Context{Context: context.Background()}
	sys := &typgo.BuildSys{Descriptor: &typgo.Descriptor{}}

	testPrj := &typgo.GoTest{}
	command := testPrj.Command(sys)

	require.Equal(t, "test", command.Name)
	require.Equal(t, "Test the project", command.Usage)
	require.Equal(t, []string{"t"}, command.Aliases)
	require.NoError(t, command.Action(c))
}

func TestTestProject_NoProjectLayout(t *testing.T) {
	testProj := &typgo.GoTest{}
	c := &typgo.Context{
		Context:  &cli.Context{Context: context.Background()},
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
	}

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	require.NoError(t, testProj.Execute(c))
}

func TestTestProject_Predefined(t *testing.T) {
	defer monkey.Patch(filepath.Walk,
		func(root string, walkFn filepath.WalkFunc) error {
			walkFn("pkg1", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("pkg2", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("pkg/service_mock", &filekit.FileInfo{IsDirField: true}, nil)
			return nil
		},
	).Unpatch()
	defer execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: "go test -cover -timeout=25s ./pkg1 ./pkg2"},
	})(t)

	c := &typgo.Context{
		Context:  &cli.Context{Context: context.Background()},
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
	}

	testProj := &typgo.GoTest{
		Args:     []string{"-timeout=25s"},
		Includes: []string{"pkg*"},
		Excludes: []string{"*_mock"},
	}
	require.NoError(t, testProj.Execute(c))
}
