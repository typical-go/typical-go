package typgo_test

import (
	"flag"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestGoTest_NoPackages(t *testing.T) {
	testProj := &typgo.GoTest{}
	c := &typgo.Context{}
	defer c.PatchBash([]*typgo.MockCommand{})(t)

	require.NoError(t, testProj.Execute(c))
}

func setupGoTest(t *testing.T) func(t *testing.T) {
	dirs := []string{
		"pkg1",
		"pkg2",
		"pkg_mock",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			t.Skip(err.Error())
		}
	}

	return func(t *testing.T) {
		os.RemoveAll("pkg1")
		os.RemoveAll("pkg2")
		os.RemoveAll("pkg")
	}
}

func TestGoTest(t *testing.T) {
	teardownTest := setupGoTest(t)
	defer teardownTest(t)

	c := &typgo.Context{
		Context: cli.NewContext(nil, &flag.FlagSet{}, nil),
	}
	defer c.PatchBash([]*typgo.MockCommand{
		{CommandLine: "go test -cover -timeout=25s ./pkg1 ./pkg2"},
	})(t)

	testProj := &typgo.GoTest{
		Timeout:  25 * time.Second,
		Includes: []string{"pkg*"},
		Excludes: []string{"*_mock"},
	}

	task := testProj.Task()
	require.Equal(t, "test", task.Name)
	require.Equal(t, []string{"t"}, task.Aliases)
	require.Equal(t, true, task.SkipFlagParsing)
	require.NoError(t, testProj.Execute(c))
}
