package typfactory_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typfactory"
)

type testcase struct {
	testName string
	typfactory.Writer
	expected string
}

func testWriter(t *testing.T, cases ...testcase) {
	for _, tt := range cases {
		t.Run(tt.testName, func(t *testing.T) {
			var debugger strings.Builder
			require.NoError(t, tt.Write(&debugger))
			require.Equal(t, tt.expected, debugger.String())
		})
	}
}
