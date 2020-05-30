package execkit_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestPrintCommand(t *testing.T) {
	var debugger strings.Builder
	execkit.PrintCommand(&execkit.Command{
		Name: "some-name",
		Args: []string{"arg-1", "arg-2"},
	}, &debugger)

	require.Equal(t, "\n$ some-name arg-1 arg-2\n", debugger.String())
}
