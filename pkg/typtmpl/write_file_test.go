package typtmpl_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typtmpl"
)

type testcase struct {
	testName string
	typtmpl.Template
	expected string
}

func testTemplate(t *testing.T, cases ...testcase) {
	for _, tt := range cases {
		t.Run(tt.testName, func(t *testing.T) {
			var debugger strings.Builder
			require.NoError(t, tt.Execute(&debugger))
			require.Equal(t, tt.expected, debugger.String())
		})
	}
}
