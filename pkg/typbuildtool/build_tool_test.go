package typbuildtool_test

import (
	"testing"

	"github.com/typical-go/typical-go/pkg/common"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestBuildSequences(t *testing.T) {
	t.Run("SHOULD implement Buildtool", func(t *testing.T) {
		var _ typcore.BuildTool = typbuildtool.BuildSequences()
	})
	t.Run("SHOULD implement Commander", func(t *testing.T) {
		var _ typbuildtool.Task = typbuildtool.BuildSequences()
	})
	t.Run("SHOULD implement Builder", func(t *testing.T) {
		var _ typbuildtool.Builder = typbuildtool.BuildSequences()
	})
	t.Run("SHOULD implement Tester", func(t *testing.T) {
		var _ typbuildtool.Tester = typbuildtool.BuildSequences()
	})
	t.Run("SHOULD implement Tester", func(t *testing.T) {
		var _ typbuildtool.Cleaner = typbuildtool.BuildSequences()
	})
	t.Run("SHOULD implement Releaser", func(t *testing.T) {
		var _ typbuildtool.Releaser = typbuildtool.BuildSequences()
	})
	t.Run("SHOULD implement Publisher", func(t *testing.T) {
		var _ typbuildtool.Publisher = typbuildtool.BuildSequences()
	})
	t.Run("SHOULD implement Preconditioner", func(t *testing.T) {
		var _ typbuildtool.Preconditioner = typbuildtool.BuildSequences()
	})
}

func TestBuildTool_Validate(t *testing.T) {
	testcases := []struct {
		*typbuildtool.TypicalBuildTool
		expectedError string
	}{
		{
			TypicalBuildTool: typbuildtool.BuildSequences(typbuildtool.StandardBuild()),
		},
		{
			TypicalBuildTool: typbuildtool.BuildSequences(),
			expectedError:    "No build modules",
		},
		{
			TypicalBuildTool: typbuildtool.BuildSequences(common.DummyValidator("some-error")),
			expectedError:    "BuildTool: some-error",
		},
	}

	for _, tt := range testcases {
		if err := tt.Validate(); err != nil {
			require.EqualError(t, err, tt.expectedError)
		} else {
			require.NoError(t, err)
		}
	}
}
