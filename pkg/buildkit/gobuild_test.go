package buildkit_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/buildkit"
)

func TestGoBuild_Args(t *testing.T) {
	testcases := []struct {
		*buildkit.GoBuild
		expected []string
	}{
		{
			GoBuild:  buildkit.NewGoBuild("some-output", "some-sources"),
			expected: []string{"build", "-o", "some-output", "some-sources"},
		},
		{
			GoBuild: buildkit.NewGoBuild("some-output", "some-sources").
				SetVariable("name1", "value1").
				SetVariable("name2", "value3"),
			expected: []string{"build", "-ldflags", "-X name1=value1 -X name2=value3", "-o", "some-output", "some-sources"},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.Args())
	}

}
