package typannot_test

import (
	"go/ast"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
)

func TestTypeDecl(t *testing.T) {
	typeDecl := &typannot.TypeDecl{
		Name: "some-name",
		Docs: []string{"doc1", "doc2"},
	}
	require.Equal(t, "some-name", typeDecl.GetName())
	require.Equal(t, []string{"doc1", "doc2"}, typeDecl.GetDocs())
}

func TestStructTag(t *testing.T) {
	testcases := []struct {
		TestName string
		Tag      *ast.BasicLit
		Expected reflect.StructTag
	}{
		{
			Tag:      &ast.BasicLit{Value: "``"},
			Expected: reflect.StructTag(""),
		},
		{
			Tag:      &ast.BasicLit{Value: "`key1=value1 key2=value2`"},
			Expected: reflect.StructTag("key1=value1 key2=value2"),
		},
		{
			Tag:      &ast.BasicLit{},
			Expected: reflect.StructTag(""),
		},
		{
			Tag:      &ast.BasicLit{Value: "`"},
			Expected: reflect.StructTag(""),
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, typannot.StructTag(tt.Tag))
		})
	}
}
