package typbuildtool_test

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typbuildtool/walker"
)

func TestAutomock(t *testing.T) {
	testcases := []struct {
		typbuildtool.Automocks
		e         *walker.TypeSpecEvent
		automocks []string
	}{
		{
			e: &walker.TypeSpecEvent{
				Filename: "filename.go",
				TypeSpec: &ast.TypeSpec{},
			},
		},
		{
			e: &walker.TypeSpecEvent{
				Filename: "filename.go",
				TypeSpec: &ast.TypeSpec{Type: &ast.InterfaceType{}},
				GenDecl:  &ast.GenDecl{Doc: astComment("some doc")},
			},
			automocks: []string{"filename.go"},
		},
		{
			e: &walker.TypeSpecEvent{
				Filename: "filename.go",
				TypeSpec: &ast.TypeSpec{Type: &ast.InterfaceType{}},
				GenDecl:  &ast.GenDecl{Doc: astComment("some doc [nomock]")},
			},
		},
	}
	for _, tt := range testcases {
		require.NoError(t, tt.OnTypeSpec(tt.e))
		require.EqualValues(t, tt.automocks, tt.Automocks)
	}
}
