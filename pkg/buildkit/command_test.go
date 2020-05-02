package buildkit_test

import (
	"context"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/buildkit"
)

func TestCommand(t *testing.T) {
	debugger := new(strings.Builder)
	input := strings.NewReader("hello world")
	ctx := context.Background()

	cmd := &buildkit.Command{
		Name:   "noname",
		Args:   []string{"arg1", "arg2", "arg3"},
		Stdout: debugger,
		Stderr: debugger,
		Stdin:  input,
		Dir:    "some-dir",
	}

	expected := exec.CommandContext(ctx, "noname", []string{"arg1", "arg2", "arg3"}...)
	expected.Stdout = debugger
	expected.Stderr = debugger
	expected.Stdin = input
	expected.Dir = "some-dir"

	require.Equal(t, "noname arg1 arg2 arg3", cmd.String())
	require.Equal(t, expected, cmd.ExecCmd(ctx))

	cmd.Print(debugger)
	require.Equal(t, "\n$ noname arg1 arg2 arg3\n", debugger.String())
}