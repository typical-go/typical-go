package typgo_test

import (
	"context"
	"errors"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestBash(t *testing.T) {
	var out strings.Builder
	input := strings.NewReader("hello world")
	ctx := context.Background()

	cmd := &typgo.Bash{
		Name:   "noname",
		Args:   []string{"arg1", "arg2", "arg3"},
		Stdout: &out,
		Stderr: &out,
		Stdin:  input,
		Dir:    "some-dir",
	}

	expected := exec.CommandContext(ctx, "noname", []string{"arg1", "arg2", "arg3"}...)
	expected.Stdout = &out
	expected.Stderr = &out
	expected.Stdin = input
	expected.Dir = "some-dir"

	require.Equal(t, expected, cmd.ExecCmd(ctx))
	require.Equal(t, cmd, cmd.Bash())
}

func TestBash_String(t *testing.T) {
	testcases := []struct {
		TestName string
		typgo.Bash
		Expected string
	}{
		{
			Bash:     typgo.Bash{Name: "name", Args: []string{"arg1", "arg2"}},
			Expected: "name arg1 arg2",
		},
		{
			Bash: typgo.Bash{
				Name: "go",
				Args: []string{"build", "-ldflags", "-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=typical-go"},
			},
			Expected: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=typical-go\"",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.String())
		})
	}
}

func TestBash_Execute(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "name1 arg1", ReturnError: errors.New("some-error")},
	})(t)

	bash := &typgo.Bash{Name: "name1", Args: []string{"arg1"}}
	c := &typgo.Context{}
	require.EqualError(t, bash.Execute(c), "some-error")
}

func TestPatch(t *testing.T) {

	defer typgo.PatchBash([]*typgo.RunExpectation{
		{
			CommandLine: "name1 arg1",
			OutputBytes: []byte("some-output-bytes"),
			ErrorBytes:  []byte("some-error-bytes"),
			ReturnError: errors.New("some-error-1"),
		},
	})(t)

	var stdout strings.Builder
	var stderr strings.Builder

	err := typgo.RunBash(nil, &typgo.Bash{
		Name:   "name1",
		Args:   []string{"arg1"},
		Stdout: &stdout,
		Stderr: &stderr,
	})
	require.EqualError(t, err, "some-error-1")

	require.Equal(t, "some-output-bytes", stdout.String())
	require.Equal(t, "some-error-bytes", stderr.String())
}

func TestPatch_MultipleExpectation(t *testing.T) {

	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "name1 arg1", ReturnError: errors.New("some-error-1")},
		{CommandLine: "name2 arg2", ReturnError: errors.New("some-error-2")},
	})(t)

	require.EqualError(t,
		typgo.RunBash(context.Background(), &typgo.Bash{Name: "name1", Args: []string{"arg1"}}),
		"some-error-1",
	)
	require.EqualError(t,
		typgo.RunBash(context.Background(), &typgo.Bash{Name: "name2", Args: []string{"arg2"}}),
		"some-error-2",
	)
}

func TestPatch_CommandNoMatchedByName(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "name1 arg2"},
	})(t)

	err := typgo.RunBash(context.Background(), &typgo.Bash{Name: "wrong", Args: []string{"arg2"}})
	require.EqualError(t, err, "typgo-mock: \"wrong arg2\" should be \"name1 arg2\"")
}

func TestPatch_CommandNoMatchedByArgs(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "name2 arg1 arg2"},
	})(t)

	err := typgo.RunBash(context.Background(), &typgo.Bash{Name: "name2", Args: []string{"arg2"}})
	require.EqualError(t, err, "typgo-mock: \"name2 arg2\" should be \"name2 arg1 arg2\"")
}

func TestPatch_NoRunExpectation(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{})(t)
	ctx := context.Background()
	err := typgo.RunBash(ctx, &typgo.Bash{Name: "name1", Args: []string{"arg1"}})
	require.EqualError(t, err, "typgo-mock: no run expectation for \"name1 arg1\"")
}
