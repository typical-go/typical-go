package app_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/typical-go/typical-go/pkg/oskit"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/internal/app"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func TestRun(t *testing.T) {
	var out strings.Builder
	defer oskit.PatchStdout(&out)()
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "go build -ldflags \"-X github.com/typical-go/typical-go/pkg/typgo.ProjectName=some-pkg -X github.com/typical-go/typical-go/pkg/typgo.ProjectPkg=some-pkg -X github.com/typical-go/typical-go/pkg/typgo.TypicalTmp=.typical-tmp\" -o .typical-tmp/bin/typical-build ./tools/typical-build"},
		{CommandLine: ".typical-tmp/bin/typical-build"},
	})(t)
	defer oskit.MkdirAll(".typical-tmp")()

	require.NoError(t, app.Run(cliContext([]string{
		"-project-pkg=some-pkg",
	})))
	require.Equal(t, "Build tools/typical-build to .typical-tmp/bin/typical-build\n", out.String())
}

func TestRun_GetParamError(t *testing.T) {
	defer typgo.PatchBash([]*typgo.RunExpectation{
		{CommandLine: "go list -m", ReturnError: errors.New("some-error")},
	})(t)
	defer oskit.MkdirAll(".typical-tmp")()

	err := app.Run(cliContext([]string{}))
	require.EqualError(t, err, "some-error: ")
}
