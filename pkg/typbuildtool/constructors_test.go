package typbuildtool_test

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

func TestConstructors(t *testing.T) {
	testcases := []struct {
		typbuildtool.Autowires
		event    *walker.AnnotationEvent
		autowire []string
	}{
		{
			event: &walker.AnnotationEvent{
				Annotation: &walker.Annotation{
					Name: "autowire",
				},
				DeclEvent: &walker.DeclEvent{
					SourceName: "SomeFunction",
					File:       &ast.File{Name: &ast.Ident{Name: "pkg"}},
				},
			},
			autowire: []string{"pkg.SomeFunction"},
		},
	}
	for _, tt := range testcases {
		require.NoError(t, tt.OnAnnotation(tt.event))
		require.EqualValues(t, tt.autowire, tt.Autowires)
	}
}
