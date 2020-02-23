package buildkit_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/buildkit"
)

func TestGoBuild(t *testing.T) {
	testcases := []struct {
		*buildkit.GoBuild
		expected []string
	}{
		{
			GoBuild:  buildkit.NewGoBuild("some-output", "some-sources"),
			expected: []string{"go", "build", "-o", "some-output", "some-sources"},
		},
		{
			GoBuild: func() *buildkit.GoBuild {
				gobuild := buildkit.NewGoBuild("some-output", "some-sources")
				gobuild.SetVariable("name1", "value1")
				gobuild.SetVariable("name2", "value3")
				return gobuild
			}(),
			expected: []string{"go", "build", "-ldflags", "-X name1=value1 -X name2=value3", "-o", "some-output", "some-sources"},
		},
	}

	for _, tt := range testcases {
		require.Equal(t, tt.expected, tt.Command(context.Background()).Args)
	}

}
