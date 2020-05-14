package typgo_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestBuildPhase(t *testing.T) {
	require.Equal(t, "test_phase", typgo.TestPhase.String())
	require.Equal(t, "compile_phase", typgo.CompilePhase.String())
	require.Equal(t, "run_phase", typgo.RunPhase.String())
	require.Equal(t, "release_phase", typgo.ReleasePhase.String())
	require.Equal(t, "publish_phase", typgo.PublishPhase.String())
	require.Equal(t, "clean_phase", typgo.CleanPhase.String())
}
