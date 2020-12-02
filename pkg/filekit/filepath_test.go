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
			walkFn("pkg1", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("pkg2", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("pkg/service_mock", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("abc1", &filekit.FileInfo{IsDirField: true}, nil)
			walkFn("abc2", &filekit.FileInfo{IsDirField: true}, nil)
			return nil
		},
	).Unpatch()

	dirs, err := filekit.FindDir([]string{"pkg*"}, []string{"**/*_mock"})
	require.NoError(t, err)
	require.Equal(t, []string{"./pkg1", "./pkg2"}, dirs)
}
