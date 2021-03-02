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
	c := &typgo.Context{}
	defer c.PatchBash([]*typgo.MockBash{
		{CommandLine: "go build -o .typical-tmp/bin/goimports golang.org/x/tools/cmd/goimports"},
		{CommandLine: ".typical-tmp/bin/goimports -w some-target"},
	})(t)

	require.NoError(t, typgo.GoImports(c, "some-target"))
}

func TestGoImport_InstallToolError(t *testing.T) {
	typgo.TypicalTmp = ".typical-tmp"
	defer func() { typgo.TypicalTmp = "" }()
	c := &typgo.Context{}
	defer c.PatchBash([]*typgo.MockBash{
		{
			CommandLine: "go build -o .typical-tmp/bin/goimports golang.org/x/tools/cmd/goimports",
			ReturnError: errors.New("some-error"),
		},
	})(t)
	require.EqualError(t, typgo.GoImports(c, "some-target"), "some-error")
}
