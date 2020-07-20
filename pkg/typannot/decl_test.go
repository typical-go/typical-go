package typannot_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typannot"
)

func TestDeclType(t *testing.T) {
	var empty typannot.DeclType
	require.Equal(t, "", empty.String())
	require.Equal(t, "Function", typannot.FuncType.String())
	require.Equal(t, "Interface", typannot.InterfaceType.String())
	require.Equal(t, "Struct", typannot.StructType.String())
	require.Equal(t, "Generic", typannot.GenericType.String())
}
