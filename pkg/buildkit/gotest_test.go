package buildkit_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/buildkit"
)

func TestGotTest(t *testing.T) {
	testcases := []struct {
		*buildkit.GoTest
		expected []string
	}{
		{
			GoTest:   buildkit.NewGoTest("target1", "target2"),
			expected: []string{"go", "test", "target1", "target2"},
		},
		{
			GoTest: buildkit.NewGoTest("target1", "target2").
				WithRace(true).WithCoverProfile("some-coverprofile"),
			expected: []string{"go", "test", "-coverprofile=some-coverprofile", "-race", "target1", "target2"},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.Command(context.Background()).Args)
	}
}
