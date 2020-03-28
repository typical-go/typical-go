package typfactory_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typfactory"
)

func TestBuildToolMain(t *testing.T) {
	testcases := []testcase{
		{
			Writer: &typfactory.BuildToolMain{DescPkg: "some-package"},
			expected: `package main

import (
	"github.com/typical-go/typical-go/pkg/typcore"
	"some-package"
)

func main() {
	typcore.LaunchBuildTool(&typical.Descriptor)
}
`,
		},
	}
	for _, tt := range testcases {
		var debugger strings.Builder
		require.NoError(t, tt.Write(&debugger))
		require.Equal(t, tt.expected, debugger.String())
	}
}
