package typast_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typast"
)

func TestDeclType(t *testing.T) {
	require.Equal(t, "Function", typast.Function.String())
	require.Equal(t, "Interface", typast.Interface.String())
	require.Equal(t, "Struct", typast.Struct.String())
	require.Equal(t, "Generic", typast.Generic.String())
}
