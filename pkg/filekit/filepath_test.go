package filekit_test

import (
	"path/filepath"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/filekit"
)

func TestMatchMulti(t *testing.T) {
	testcases := []struct {
		testName string
		patterns []string
		name     string
		expected bool
	}{
		{
			patterns: []string{"aa*", "bb*"},
			name:     "aacc",
			expected: true,
		},
		{
			patterns: []string{"aa*", "bb*"},
			name:     "bbcc",
			expected: true,
		},
		{
			patterns: []string{"aa*", "bb*"},
			name:     "abab",
			expected: false,
		},
	}
	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, filekit.MatchMulti(tt.patterns, tt.name))
		})
	}
}

func TestFindDir(t *testing.T) {
	defer monkey.Patch(filepath.Walk,
		func(root string, walkFn filepath.WalkFunc) error {
			dirs := []string{
				"internal/app/infra",
				"internal/app/domain",
				"internal/app/domain/mylibrary",
				"internal/app/domain/mylibrary/controller",
				"internal/app/domain/mylibrary/service",
				"internal/app/domain/mylibrary/repo",
				"internal/app/domain/mybook/controller",
				"internal/app/domain/mybook/service",
				"internal/app/domain/mybook/repo",
				"internal/app/generated",
				"internal/app/generated/constructor",
				"internal/app/generated/config",
			}

			for _, dir := range dirs {
				walkFn(dir, &filekit.FileInfo{IsDirField: true}, nil)
			}

			return nil
		},
	).Unpatch()

	dirs, err := filekit.FindDir(
		[]string{"internal/**/*"},
		[]string{"internal/**/generated/**"},
	)
	require.NoError(t, err)
	require.Equal(t, []string{
		"./internal/app/infra",
		"./internal/app/domain",
		"./internal/app/domain/mylibrary",
		"./internal/app/domain/mylibrary/controller",
		"./internal/app/domain/mylibrary/service",
		"./internal/app/domain/mylibrary/repo",
		"./internal/app/domain/mybook/controller",
		"./internal/app/domain/mybook/service",
		"./internal/app/domain/mybook/repo",
		"./internal/app/generated",
	}, dirs)
}
