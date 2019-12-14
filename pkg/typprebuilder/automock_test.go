package typprebuilder_test

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/pkg/typprebuilder"
	"github.com/typical-go/typical-go/pkg/typprebuilder/walker"
)

func TestAutomock(t *testing.T) {
	testcases := []struct {
		typprebuilder.Automocks
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
				TypeSpec: createInterfaceSpecWithComment("some doc"),
			},
			automocks: []string{"filename.go"},
		},
		{
			e: &walker.TypeSpecEvent{
				Filename: "filename.go",
				TypeSpec: createInterfaceSpecWithComment("some doc [nomock]"),
			},
		},
	}
	for _, tt := range testcases {
		require.NoError(t, tt.OnTypeSpec(tt.e))
		require.EqualValues(t, tt.automocks, tt.Automocks)
	}
}

func createInterfaceSpecWithComment(comment string) *ast.TypeSpec {
	return &ast.TypeSpec{
		Type: &ast.InterfaceType{},
		Doc:  astComment(comment),
	}
}
