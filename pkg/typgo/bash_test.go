package typgo_test

import (
	"context"
	"os"
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

func TestBashCommand(t *testing.T) {
	testcases := []struct {
		TestName     string
		Line         string
		ExpectedBash *typgo.Bash
	}{
		{
			Line: "go build -o output",
			ExpectedBash: &typgo.Bash{
				Name:   "go",
				Args:   []string{"build", "-o", "output"},
				Stdout: os.Stdout,
				Stderr: os.Stderr,
				Stdin:  os.Stdin,
			},
		},
		{
			Line: "dir",
			ExpectedBash: &typgo.Bash{
				Name:   "dir",
				Args:   []string{},
				Stdout: os.Stdout,
				Stderr: os.Stderr,
				Stdin:  os.Stdin,
			},
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.ExpectedBash, typgo.BashCommand(tt.Line))
		})
	}

}
