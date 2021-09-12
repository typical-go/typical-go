package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestImportAliases(t *testing.T) {
	i := typgen.NewAliasGenerator(nil)

	require.Equal(t, "a", i.Generate("package1"))
	require.Equal(t, "b", i.Generate("package2"))

	require.EqualValues(t, map[string]string{
		"package1": "a",
		"package2": "b",
	}, i.Map)

	require.EqualValues(t, []*typgen.Import{
		{Name: "a", Path: "package1"},
		{Name: "b", Path: "package2"},
	}, i.Imports())
}
