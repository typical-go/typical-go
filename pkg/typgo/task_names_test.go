package typgo_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/oskit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func TestBuildCmdRuns(t *testing.T) {
	var out strings.Builder
	defer oskit.PatchStdout(&out)()
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "./typicalw run-1", OutputBytes: []byte("Running (1)\n")},
		{CommandLine: "./typicalw run-2", OutputBytes: []byte("Running (2)\n")},
		{CommandLine: "./typicalw run-3", OutputBytes: []byte("Running (3)\n")},
	})(t)
	c := &typgo.Context{Context: cli.NewContext(nil, nil, nil)}
	typgo.ProjectName = "proj01"

	sr := typgo.TaskNames{"run-1", "run-2", "run-3"}
	require.NoError(t, sr.Execute(c))
	require.Equal(t, "\n----- proj01: ./typicalw run-1 -----\n\nRunning (1)\n\n----- proj01: ./typicalw run-2 -----\n\nRunning (2)\n\n----- proj01: ./typicalw run-3 -----\n\nRunning (3)\n", out.String())
}

func TestBuildCmdRuns_Error(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "./typicalw run-1"},
		{CommandLine: "./typicalw run-2", ReturnError: errors.New("some-error")},
	})(t)

	c := &typgo.Context{Context: cli.NewContext(nil, nil, nil)}
	typgo.ProjectName = "proj01"

	sr := typgo.TaskNames{"run-1", "run-2", "run-3"}
	require.EqualError(t, sr.Execute(c), "some-error")
}
