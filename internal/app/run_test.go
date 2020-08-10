package app_test

import (
	"errors"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/internal/app"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestRun(t *testing.T) {
	var output strings.Builder
	app.Stdout = &output
	defer func() { app.Stdout = os.Stdout }()

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{
			"go", "build",
			"-ldflags", "-X github.com/typical-go/typical-go/pkg/typgo.ProjectPkg=some-pkg -X github.com/typical-go/typical-go/pkg/typgo.TypicalTmp=.typical-tmp",
			"-o", ".typical-tmp/bin/typical-build",
			"./tools/typical-build",
		}},
		{CommandLine: []string{".typical-tmp/bin/typical-build"}},
	})
	defer unpatch(t)

	os.Mkdir(".typical-tmp", 0777)
	defer os.RemoveAll(".typical-tmp")

	err := app.Run(cliContext([]string{
		"-project-pkg=some-pkg",
	}))
	require.NoError(t, err)
	require.Equal(t, "Build tools/typical-build to .typical-tmp/bin/typical-build\n", output.String())
}

func TestRun_GetParamError(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"go", "list", "-m"}, ReturnError: errors.New("some-error")},
	})
	defer unpatch(t)

	os.Mkdir(".typical-tmp", 0777)
	defer os.RemoveAll(".typical-tmp")

	err := app.Run(cliContext([]string{}))
	require.EqualError(t, err, "some-error: ")
}
