package buildkit_test

import (
	"testing"
	"time"

	"github.com/typical-go/typical-go/pkg/buildkit"

	"github.com/stretchr/testify/require"
)

func TestGotTest_Args(t *testing.T) {
	testcases := []struct {
		*buildkit.GoTest
		expected string
	}{
		{
			GoTest: &buildkit.GoTest{
				Targets: []string{"target1", "target2"},
			},
			expected: "go test target1 target2",
		},
		{
			GoTest: &buildkit.GoTest{
				Targets:      []string{"target1", "target2"},
				Timeout:      10 * time.Second,
				Race:         true,
				CoverProfile: "some-coverprofile",
			},

			expected: "go test -timeout=10s -coverprofile=some-coverprofile -race target1 target2",
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.Command().String())
	}
}
