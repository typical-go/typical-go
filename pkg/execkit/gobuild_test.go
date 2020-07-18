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
		expected *execkit.Command
	}{
		{
			GoBuild: &execkit.GoBuild{
				Output:      "some-output",
				MainPackage: "some-sources",
			},
			expected: &execkit.Command{
				Name: "go",
				Args: []string{"build", "-o", "some-output", "some-sources"},
			},
		},
		{
			GoBuild: &execkit.GoBuild{
				Output:      "some-output",
				MainPackage: "some-sources",
				Ldflags: execkit.BuildVars{
					"name1": "value1",
					"name2": "value3",
				},
			},
			expected: &execkit.Command{
				Name: "go",
				Args: []string{
					"build",
					"-ldflags", "-X name1=value1 -X name2=value3",
					"-o", "some-output",
					"some-sources",
				},
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.testName, func(t *testing.T) {
			cmd := tt.Command()
			require.Equal(t, tt.expected.Name, cmd.Name)
			require.Equal(t, tt.expected.Args, cmd.Args)
		})
	}
}
