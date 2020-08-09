package app_test

import (
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
			"-o", "some-dir/.typical-tmp/bin/typical-build",
			"./tools/typical-build",
		}},
		{CommandLine: []string{"some-dir/.typical-tmp/bin/typical-build"}},
	})
	defer unpatch(t)

	os.MkdirAll("some-dir/.typical-tmp", 0777)
	defer os.RemoveAll("some-dir")

	err := app.Run(cliContext([]string{
		"-project-dir=some-dir",
		"-project-pkg=some-pkg",
	}))
	require.NoError(t, err)
	require.Equal(t, "Build tools/typical-build to some-dir/.typical-tmp/bin/typical-build\n", output.String())
}
