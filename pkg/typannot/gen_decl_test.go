package typannot_test

import (
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
