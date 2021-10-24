package typgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgen"
)

func TestWriteSourceCode(t *testing.T) {
	testcases := []struct {
		TestName string
		Coder    typgen.Coder
		Expected string
	}{
		{
			Coder: typgen.Coders{
				typgen.CodeLine("some-code-1"),
				typgen.CodeLine("some-code-2"),
			},
			Expected: "some-code-1\nsome-code-2\n",
		},
		{
			Coder: typgen.CodeLines{
				"some-code-1",
				"some-code-2",
			},
			Expected: "some-code-1\nsome-code-2\n",
		},
	}
	for _, tt := range testcases {
		t.Run(tt.TestName, func(t *testing.T) {
			require.Equal(t, tt.Expected, tt.Coder.Code())
		})
	}
}
