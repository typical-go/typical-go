package typgo_test

import (
	"flag"
	"path/filepath"
	"testing"
	"time"

	"github.com/typical-go/typical-go/pkg/filekit"

	"bou.ke/monkey"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestGoTest_NoPackages(t *testing.T) {
	testProj := &typgo.GoTest{}
	c := &typgo.Context{}
	defer c.PatchBash([]*typgo.MockBash{})(t)

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

	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
	}
	defer c.PatchBash([]*typgo.MockBash{
		{CommandLine: "go test -cover -timeout=25s ./pkg1 ./pkg2"},
	})(t)

	testProj := &typgo.GoTest{
		Timeout:  25 * time.Second,
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
