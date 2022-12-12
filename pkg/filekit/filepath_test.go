package filekit_test

import (
	"os"
	"testing"

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

func setupFindDirTest(t *testing.T) func(t *testing.T) {
	dirs := []string{
		"tmp/internal/app/infra",
		"tmp/internal/app/domain/mylibrary/controller",
		"tmp/internal/app/domain/mylibrary/repo",
		"tmp/internal/app/domain/mylibrary/service",
		"tmp/internal/app/domain/mybook/controller",
		"tmp/internal/app/domain/mybook/repo",
		"tmp/internal/app/domain/mybook/service",
		"tmp/internal/app/generated/constructor",
		"tmp/internal/app/generated/config",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil && !os.IsExist(err) {
			t.Skip(err.Error())
		}
	}

	return func(t *testing.T) {
		os.RemoveAll("tmp")
	}
}

func TestFindDir(t *testing.T) {
	teardownTest := setupFindDirTest(t)
	defer teardownTest(t)

	dirs, err := filekit.FindDir(
		[]string{"tmp/internal/**/*"},
		[]string{"tmp/internal/**/generated/**"},
	)
	require.NoError(t, err)
	require.Equal(t, []string{
		"./tmp/internal/app",
		"./tmp/internal/app/domain",
		"./tmp/internal/app/domain/mybook",
		"./tmp/internal/app/domain/mybook/controller",
		"./tmp/internal/app/domain/mybook/repo",
		"./tmp/internal/app/domain/mybook/service",
		"./tmp/internal/app/domain/mylibrary",
		"./tmp/internal/app/domain/mylibrary/controller",
		"./tmp/internal/app/domain/mylibrary/repo",
		"./tmp/internal/app/domain/mylibrary/service",
		"./tmp/internal/app/generated",
		"./tmp/internal/app/infra",
	}, dirs)
}
