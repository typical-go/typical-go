package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestDeclType(t *testing.T) {
	var empty typast.DeclType
	require.Equal(t, "", empty.String())
	require.Equal(t, "Function", typast.FuncType.String())
	require.Equal(t, "Interface", typast.InterfaceType.String())
	require.Equal(t, "Struct", typast.StructType.String())
	require.Equal(t, "Generic", typast.GenericType.String())
}
