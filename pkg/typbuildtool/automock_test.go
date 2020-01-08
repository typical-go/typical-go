package typbuildtool_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

func TestAutomock(t *testing.T) {
	testcases := []struct {
		typbuildtool.Automocks
		e         *walker.DeclEvent
		automocks []string
	}{
		{
			e: &walker.DeclEvent{
				Filename: "filename.go",
			},
		},
		{
			e: &walker.DeclEvent{
				Filename:  "filename.go",
				EventType: walker.InterfaceType,
				Doc:       "some doc [mock]",
			},
			automocks: []string{"filename.go"},
		},
		{
			e: &walker.DeclEvent{
				Filename:  "filename.go",
				EventType: walker.InterfaceType,
				Doc:       "some doc",
			},
		},
	}
	for _, tt := range testcases {
		require.NoError(t, tt.OnDecl(tt.e))
		require.EqualValues(t, tt.automocks, tt.Automocks)
	}
}
