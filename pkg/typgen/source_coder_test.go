package typgen_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestWriteSourceCode(t *testing.T) {
	filename := "test_WriteSourceCode"
	defer os.Remove(filename)
	typgen.WriteSourceCode(filename,
		typgen.Comment("some comment 1"),
		typgen.Comment("some comment 2"),
	)

	b, _ := ioutil.ReadFile(filename)
	require.Equal(t, "// some comment 1\n// some comment 2\n", string(b))
}
