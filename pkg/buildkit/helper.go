package buildkit

import (
	"context"
	"os/exec"
	"strings"
)

// AvailableCommand return path and boolean whether the command not available in bash
func AvailableCommand(ctx context.Context, name string) (ok bool) {
	if name == "" {
		return false
	}

	var debugger strings.Builder
	cmd := exec.CommandContext(ctx, "command", "-v", name)
	cmd.Stdout = &debugger

	if err := cmd.Run(); err != nil {
		return false
	}

	return strings.TrimSpace(debugger.String()) != ""
}
