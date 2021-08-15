package typgo_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestPatch(t *testing.T) {
	mocker := &typgo.MockCommandRunner{
		Mocks: []*typgo.MockCommand{
			{CommandLine: "name1 arg1", OutputBytes: []byte("some-output-bytes"), ErrorBytes: []byte("some-error-bytes"), ReturnError: errors.New("some-error-1")},
		},
	}

	var stdout strings.Builder
	var stderr strings.Builder

	err := mocker.Run(&typgo.Command{
		Name:   "name1",
		Args:   []string{"arg1"},
		Stdout: &stdout,
		Stderr: &stderr,
	})
	require.EqualError(t, err, "some-error-1")

	require.Equal(t, "some-output-bytes", stdout.String())
	require.Equal(t, "some-error-bytes", stderr.String())
	require.NoError(t, mocker.Close())
}

func TestPatch_MultipleExpectation(t *testing.T) {
	mocker := &typgo.MockCommandRunner{
		Mocks: []*typgo.MockCommand{
			{CommandLine: "name1 arg1", ReturnError: errors.New("some-error-1")},
			{CommandLine: "name2 arg2", ReturnError: errors.New("some-error-2")},
		},
	}

	require.EqualError(t, mocker.Run(&typgo.Command{Name: "name1", Args: []string{"arg1"}}), "some-error-1")
	require.EqualError(t, mocker.Run(&typgo.Command{Name: "name2", Args: []string{"arg2"}}), "some-error-2")
	require.NoError(t, mocker.Close())
}
func TestPatch_MissingCall(t *testing.T) {
	mocker := &typgo.MockCommandRunner{
		Mocks: []*typgo.MockCommand{
			{CommandLine: "name1 arg1", ReturnError: errors.New("some-error-1")},
			{CommandLine: "name2 arg2", ReturnError: errors.New("some-error-2")},
		},
	}

	require.EqualError(t, mocker.Run(&typgo.Command{Name: "name1", Args: []string{"arg1"}}), "some-error-1")
	require.EqualError(t, mocker.Close(), "missing bash call: \"name2 arg2\"")
}

func TestPatch_CommandNoMatchedByName(t *testing.T) {
	mocker := &typgo.MockCommandRunner{
		Mocks: []*typgo.MockCommand{
			{CommandLine: "name1 arg2"},
		},
	}

	err := mocker.Run(&typgo.Command{Name: "wrong", Args: []string{"\"arg2\""}})
	require.EqualError(t, err, "typgo-mock: \"wrong \"arg2\"\" should be \"name1 arg2\"")
	require.NoError(t, mocker.Close())
}

func TestPatch_CommandNoMatchedByArgs(t *testing.T) {
	mocker := &typgo.MockCommandRunner{
		Mocks: []*typgo.MockCommand{
			{CommandLine: "name2 arg1 arg2"},
		},
	}

	err := mocker.Run(&typgo.Command{Name: "name2", Args: []string{"arg2"}})
	require.EqualError(t, err, "typgo-mock: \"name2 arg2\" should be \"name2 arg1 arg2\"")
	require.NoError(t, mocker.Close())
}

func TestPatch_NoRunExpectation(t *testing.T) {
	mocker := &typgo.MockCommandRunner{}
	err := mocker.Run(&typgo.Command{Name: "name1", Args: []string{"arg1"}})
	require.EqualError(t, err, "typgo-mock: no run expectation for \"name1 arg1\"")
	require.NoError(t, mocker.Close())
}
