package exor_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/exor"

	"github.com/stretchr/testify/require"
)

func TestGoTest(t *testing.T) {
	t.Run("SHOULD implement Executor", func(t *testing.T) {
		var _ exor.Executor = exor.NewGoTest("", "")
	})
}

func TestGotTest_Args(t *testing.T) {
	testcases := []struct {
		*exor.GoTest
		expected []string
	}{
		{
			GoTest:   exor.NewGoTest("target1", "target2"),
			expected: []string{"test", "target1", "target2"},
		},
		{
			GoTest: exor.NewGoTest("target1", "target2").
				WithRace(true).WithCoverProfile("some-coverprofile"),
			expected: []string{"test", "-coverprofile=some-coverprofile", "-race", "target1", "target2"},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.Args())
	}
}
