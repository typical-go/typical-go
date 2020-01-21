package typbuildtool_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore/walker"
)

func TestAutomock(t *testing.T) {
	testcases := []struct {
		typbuildtool.Automocks
		e         *walker.AnnotationEvent
		automocks []string
	}{
		{
			e: &walker.AnnotationEvent{
				Annotation: &walker.Annotation{
					Name: "mock",
				},
				DeclEvent: &walker.DeclEvent{
					Filename: "filename.go",
				},
			},
			automocks: []string{"filename.go"},
		},
	}
	for i, tt := range testcases {
		require.NoError(t, tt.OnAnnotation(tt.e), i)
		require.EqualValues(t, tt.automocks, tt.Automocks, i)
	}
}
