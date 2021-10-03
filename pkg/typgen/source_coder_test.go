package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestWriteSourceCode(t *testing.T) {
	sourceCoders := typgen.SourceCoders{
		typgen.Comment("some comment 1"),
		typgen.Comment("some comment 2"),
	}
	require.Equal(t, "// some comment 1\n// some comment 2\n", sourceCoders.SourceCode())
}
