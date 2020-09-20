package execkit_test

import (
	"context"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestCommand(t *testing.T) {
	var out strings.Builder
	input := strings.NewReader("hello world")
	ctx := context.Background()

	cmd := &execkit.Command{
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
	require.Equal(t, cmd, cmd.Command())
}

func TestCommand_String(t *testing.T) {
	testcases := []struct {
		TestName string
		execkit.Command
		Expected string
	}{
		{
			Command:  execkit.Command{Name: "name", Args: []string{"arg1", "arg2"}},
			Expected: "name arg1 arg2",
		},
		{
			Command: execkit.Command{
				Name: "go",
				Args: []string{"build", "-ldflags", "-X github.com/typical-go/typical-go/pkg/typgo.AppName=typical-go"},
			},
			Expected: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.AppName=typical-go\"",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.String())
		})
	}
}
