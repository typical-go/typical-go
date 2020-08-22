package typannot_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
)

func TestFuncDecl(t *testing.T) {
	funcDecl := &typannot.FuncDecl{
		Name: "some-name",
		Docs: []string{"doc1", "doc2"},
	}
	require.Equal(t, "some-name", funcDecl.GetName())
	require.Equal(t, []string{"doc1", "doc2"}, funcDecl.GetDocs())
}
