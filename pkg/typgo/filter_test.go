package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestNoPrefix(t *testing.T) {
	testcases := []struct {
		prefixes []string
		message  string
		expected bool
	}{
		{
			prefixes: []string{"revision"},
			message:  "revision: something",
			expected: true,
		},
		{
			prefixes: []string{"revision"},
			message:  "REVISION: something",
			expected: true,
		},
		{
			message: "something",
		},
	}
	for _, tt := range testcases {
		filter := typgo.ExcludePrefix(tt.prefixes...)
		require.Equal(t, tt.expected, filter.Exclude(tt.message))
	}
}
