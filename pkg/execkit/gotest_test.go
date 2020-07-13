package execkit_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestGotTest_Args(t *testing.T) {
	testcases := []struct {
		testName string
		*execkit.GoTest
		expected string
	}{
		{
			GoTest: &execkit.GoTest{
				Packages: []string{"target1", "target2"},
			},
			expected: "go test target1 target2",
		},
		{
			GoTest: &execkit.GoTest{
				Packages:     []string{"target1", "target2"},
				Timeout:      10 * time.Second,
				Race:         true,
				CoverProfile: "some-coverprofile",
			},

			expected: "go test -timeout=10s -coverprofile=some-coverprofile -race target1 target2",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.String())
		})
	}
}
