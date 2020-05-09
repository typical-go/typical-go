package github_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/github"
)

func TestNoPrefix(t *testing.T) {
	testcases := []struct {
		prefixes []string
		message  string
		expected string
	}{
		{
			prefixes: []string{"revision"},
			message:  "revision: something",
		},
		{
			prefixes: []string{"revision"},
			message:  "REVISION: something",
		},
		{
			message:  "something",
			expected: "something",
		},
	}
	for _, tt := range testcases {
		filter := github.NoPrefix(tt.prefixes...)
		require.Equal(t, tt.expected, filter.ReleaseFilter(tt.message))
	}
}
