package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestImportAliases(t *testing.T) {
	importAliases := typgen.NewImportAliases()

	require.Equal(t, "a", importAliases.Append("package1"))
	require.Equal(t, "b", importAliases.Append("package2"))

	require.Equal(t, map[string]string{
		"package1": "a",
		"package2": "b",
	}, importAliases.Map)
}
