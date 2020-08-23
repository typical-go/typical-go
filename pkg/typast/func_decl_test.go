package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestFuncDecl(t *testing.T) {
	funcDecl := &typast.FuncDecl{
		Name: "some-name",
		Docs: []string{"doc1", "doc2"},
	}
	require.Equal(t, "some-name", funcDecl.GetName())
	require.Equal(t, []string{"doc1", "doc2"}, funcDecl.GetDocs())
}

func TestFuncDecl_IsMethod(t *testing.T) {
	testnames := []struct {
		TestName string
		*typast.FuncDecl
		Expected bool
	}{
		{
			FuncDecl: &typast.FuncDecl{Recv: &typast.FieldList{}},
			Expected: true,
		},
		{
			FuncDecl: &typast.FuncDecl{},
			Expected: false,
		},
	}
	for _, tt := range testnames {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.IsMethod())
		})
	}
}
