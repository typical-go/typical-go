package buildkit_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/buildkit"
)

func TestCommandExist(t *testing.T) {
	testcases := []struct {
		name     string
		expected bool
	}{
		{"go", true},
		{"", false},
		{"invalid-command", false},
	}

	ctx := context.Background()
	for _, tt := range testcases {
		require.Equal(t, tt.expected, buildkit.AvailableCommand(ctx, tt.name))
	}
}
