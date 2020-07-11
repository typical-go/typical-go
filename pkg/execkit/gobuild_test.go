package execkit_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestGoBuild(t *testing.T) {
	testcases := []struct {
		testName string
		*execkit.GoBuild
		expected string
	}{
		{
			GoBuild: &execkit.GoBuild{
				Output: "some-output",
				Source: "some-sources",
			},
			expected: "go build -o some-output some-sources",
		},
		{
			GoBuild: &execkit.GoBuild{
				Output: "some-output",
				Source: "some-sources",
				Ldflags: execkit.BuildVars{
					"name1": "value1",
					"name2": "value3",
				},
			},
			expected: "go build -ldflags -X name1=value1 -X name2=value3 -o some-output some-sources",
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			require.Equal(t, tt.expected, tt.String())
		})
	}
}
