package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestFuncDecl(t *testing.T) {
	funcDecl := &typgen.FuncDecl{
		Name: "some-name",
		Docs: []string{"doc1", "doc2"},
	}
	require.Equal(t, "some-name", funcDecl.GetName())
	require.Equal(t, []string{"doc1", "doc2"}, funcDecl.GetDocs())
}

func TestFuncDecl_IsMethod(t *testing.T) {
	testnames := []struct {
		TestName string
		*typgen.FuncDecl
		Expected bool
	}{
		{
			FuncDecl: &typgen.FuncDecl{Recv: &typgen.FieldList{}},
			Expected: true,
		},
		{
			FuncDecl: &typgen.FuncDecl{},
			Expected: false,
		},
	}
	for _, tt := range testnames {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.IsMethod())
		})
	}
}
