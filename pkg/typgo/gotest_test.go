package typgo_test

import (
	"context"
	"flag"
	"path/filepath"
	"testing"

	"github.com/typical-go/typical-go/pkg/filekit"

	"bou.ke/monkey"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestGoTest_NoPackages(t *testing.T) {
	testProj := &typgo.GoTest{}
	c := &typgo.Context{
		Context:  &cli.Context{Context: context.Background()},
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
	}

	defer typgo.PatchBash([]*typgo.RunExpectation{})(t)

	require.NoError(t, testProj.Execute(c))
}

func TestGoTest(t *testing.T) {
	defer monkey.Patch(filepath.Walk,
		func(root string, walkFn filepath.WalkFunc) error {
			walkFn("pkg1", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("pkg2", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("pkg/service_mock", &filekit.FileInfo{IsDirField: true}, nil)
			return nil
		},
	).Unpatch()
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "go test -cover -timeout=25s ./pkg1 ./pkg2"},
	})(t)

	c := &typgo.Context{
		Context:  gotestCliContext(nil),
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
	}

	testProj := &typgo.GoTest{
		Args:     []string{"-timeout=25s"},
		Includes: []string{"pkg*"},
		Excludes: []string{"*_mock"},
	}
	require.NoError(t, testProj.Execute(c))
}

func TestGoTest_WithCoverProfile(t *testing.T) {
	defer monkey.Patch(filepath.Walk,
		func(root string, walkFn filepath.WalkFunc) error {
			walkFn("pkg1", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("pkg2", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("pkg/service_mock", &filekit.FileInfo{IsDirField: true}, nil)
			return nil
		},
	).Unpatch()
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "go test -coverprofile=cover.out -timeout=25s ./pkg1 ./pkg2"},
	})(t)

	c := &typgo.Context{
		Context:  gotestCliContext([]string{"-coverprofile=cover.out"}),
		BuildSys: &typgo.BuildSys{Descriptor: &typgo.Descriptor{}},
	}

	testProj := &typgo.GoTest{
		Args:     []string{"-timeout=25s"},
		Includes: []string{"pkg*"},
		Excludes: []string{"*_mock"},
	}
	require.NoError(t, testProj.Execute(c))
}

func gotestCliContext(args []string) *cli.Context {
	flagSet := &flag.FlagSet{}
	flagSet.String("coverprofile", "", "")
	flagSet.Parse(args)
	return cli.NewContext(nil, flagSet, nil)
}
