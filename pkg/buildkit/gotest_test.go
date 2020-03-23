package buildkit_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/buildkit"

	"github.com/stretchr/testify/require"
)

func TestGotTest_Args(t *testing.T) {
	testcases := []struct {
		*buildkit.GoTest
		expected []string
	}{
		{
			GoTest:   buildkit.NewGoTest("target1", "target2"),
			expected: []string{"test", "target1", "target2"},
		},
		{
			GoTest: buildkit.NewGoTest("target1", "target2").
				WithRace(true).WithCoverProfile("some-coverprofile"),
			expected: []string{"test", "-coverprofile=some-coverprofile", "-race", "target1", "target2"},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.Args())
	}
}
