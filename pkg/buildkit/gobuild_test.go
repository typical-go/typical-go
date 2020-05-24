package buildkit_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/buildkit"
)

func TestGoBuild(t *testing.T) {
	testcases := []struct {
		*buildkit.GoBuild
		expected string
	}{
		{
			GoBuild: &buildkit.GoBuild{
				Out:    "some-output",
				Source: "some-sources",
			},
			expected: "go build -o some-output some-sources",
		},
		{
			GoBuild: &buildkit.GoBuild{
				Out:    "some-output",
				Source: "some-sources",
				Ldflags: []string{
					buildkit.BuildVar("name1", "value1"),
					buildkit.BuildVar("name2", "value3"),
				},
			},
			expected: "go build -ldflags -X name1=value1 -X name2=value3 -o some-output some-sources",
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.Command().String())
	}

}
