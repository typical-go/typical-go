package typgo_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestGoImport(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp"
	defer func() { typgo.TypicalTmp = "" }()

	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "go build -o .typical-tmp/bin/goimports golang.org/x/tools/cmd/goimports"},
		{CommandLine: ".typical-tmp/bin/goimports -w some-target"},
	})(t)

	require.NoError(t, typgo.GoImports("some-target"))
}

func TestGoImport_InstallToolError(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp"
	defer func() { typgo.TypicalTmp = "" }()

	defer typgo.PatchBash([]*typgo.RunExpectation{
		{
			CommandLine: "go build -o .typical-tmp/bin/goimports golang.org/x/tools/cmd/goimports",
			ReturnError: errors.New("some-error"),
		},
	})(t)

	require.EqualError(t, typgo.GoImports("some-target"), "some-error")
}
