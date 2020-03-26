package buildkit_test

import (
	"context"
	"testing"
	"time"

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
			expected: []string{"test", "-timeout=20s", "target1", "target2"},
		},
		{
			GoTest: buildkit.NewGoTest("target1", "target2").
				WithTimeout(10 * time.Second).
				WithRace(true).
				WithCoverProfile("some-coverprofile"),
			expected: []string{"test", "-timeout=10s", "-coverprofile=some-coverprofile", "-race", "target1", "target2"},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.Args())
	}
}

func TestGoTest_Execute(t *testing.T) {
	t.Run("GIVEN no test targets", func(t *testing.T) {
		require.EqualError(t,
			buildkit.NewGoTest().Execute(context.Background()),
			"Nothing to test",
		)
	})
}
