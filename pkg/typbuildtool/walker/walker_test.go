package walker

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWalkTarget(t *testing.T) {
	testcases := []struct {
		filename string
		result   bool
	}{
		{"file.go", true},
		{"file_test.go", false},
		{"file.test.go", true},
		{"file", false},
	}
	for _, tt := range testcases {
		require.Equal(t, tt.result, isWalkTarget(tt.filename))
	}
}
