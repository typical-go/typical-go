package tmplkit_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/tmplkit"
)

func TestImportAliases(t *testing.T) {
	importAliases := tmplkit.NewImports(nil)

	require.Equal(t, "a", importAliases.AppendWithAlias("package1"))
	require.Equal(t, "b", importAliases.AppendWithAlias("package2"))

	require.Equal(t, map[string]string{
		"package1": "a",
		"package2": "b",
	}, importAliases.Map)
}
