package filekit_test

import (
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
