package typbuildtool_test

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

func TestAutowire(t *testing.T) {
	testcases := []struct {
		typbuildtool.Autowires
		event    *walker.DeclEvent
		autowire []string
	}{
		{
			event: &walker.DeclEvent{
				Name: "SomeFunction",
				File: &ast.File{Name: &ast.Ident{Name: "pkg"}},
			},
		},
		{
			event: &walker.DeclEvent{
				Name: "SomeFunction",
				Doc:  "some doc",
				File: &ast.File{Name: ast.NewIdent("pkg")},
			},
		},
		{
			event: &walker.DeclEvent{
				Name: "SomeFunction",
				Doc:  "some doc [autowire]",
				File: &ast.File{Name: ast.NewIdent("pkg")},
			},
			autowire: []string{"pkg.SomeFunction"},
		},
		{
			event: &walker.DeclEvent{
				Name: "SomeFunction",
				Doc:  "some doc [Autowire]",
				File: &ast.File{Name: ast.NewIdent("pkg")},
			},
			autowire: []string{"pkg.SomeFunction"},
		},
		{
			event: &walker.DeclEvent{
				Name: "SomeFunction",
				Doc:  "some doc [AUTOWIRE]",
				File: &ast.File{Name: ast.NewIdent("pkg")},
			},
			autowire: []string{"pkg.SomeFunction"},
		},
	}
	for _, tt := range testcases {
		require.NoError(t, tt.OnDecl(tt.event))
		require.EqualValues(t, tt.autowire, tt.Autowires)
	}
}
