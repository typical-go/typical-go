package app_test

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/internal/app"
	"github.com/typical-go/typical-go/pkg/execkit"
)

func TestSetup(t *testing.T) {
	var output strings.Builder
	app.Stdout = &output
	defer func() { app.Stdout = os.Stdout }()

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	os.Mkdir("somedir1", 0777)
	defer os.RemoveAll("somedir1")

	err := app.Setup(cliContext([]string{
		"-project-dir=somedir1",
		"-project-pkg=some-pkg",
	}))
	require.NoError(t, err)

	b, _ := ioutil.ReadFile("somedir1/typicalw")
	require.Equal(t, `#!/bin/bash

set -e

TYPTMP=.typical-tmp
TYPGO=$TYPTMP/bin/typical-go

if ! [ -s $TYPGO ]; then
	echo "Build typical-go"
	go build -o $TYPGO github.com/typical-go/typical-go
fi

$TYPGO run \
	-project-pkg="some-pkg" \
	-typical-build="tools/typical-build" \
	-typical-tmp=$TYPTMP \
	$@
`, string(b))

	require.Equal(t, "Create wrapper 'somedir1/typicalw'\n", output.String())
}

func TestSetup_WithGomodFlag(t *testing.T) {
	var output strings.Builder
	app.Stdout = &output
	defer func() { app.Stdout = os.Stdout }()

	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{CommandLine: []string{"go", "mod", "init", "somepkg"}},
	})
	defer unpatch(t)

	os.Mkdir("somedir2", 0777)
	defer os.RemoveAll("somedir2")

	err := app.Setup(cliContext([]string{
		"-project-dir=somedir2",
		"-project-pkg=somepkg",
		"-gomod=somepkg",
	}))
	require.NoError(t, err)

	require.Equal(t, "Initiate go.mod\nCreate wrapper 'somedir2/typicalw'\n", output.String())
}

func TestSetup_WithGomodFlag_Error(t *testing.T) {
	unpatch := execkit.Patch([]*execkit.RunExpectation{
		{
			CommandLine: []string{"go", "mod", "init", "somepkg"},
			ErrorBytes:  []byte("error-message"),
			ReturnError: errors.New("some-error"),
		},
	})
	defer unpatch(t)

	os.Mkdir("somedir2", 0777)
	defer os.RemoveAll("somedir2")

	err := app.Setup(cliContext([]string{
		"-project-dir=somedir1",
		"-project-pkg=somepkg",
		"-gomod=somepkg",
	}))
	require.EqualError(t, err, "some-error: error-message")
}

func TestSetup_WithNewFlag(t *testing.T) {
	var output strings.Builder
	app.Stdout = &output
	defer func() { app.Stdout = os.Stdout }()

	unpatch := execkit.Patch([]*execkit.RunExpectation{})
	defer unpatch(t)

	os.Mkdir("somedir4", 0777)
	defer os.RemoveAll("somedir4")

	err := app.Setup(cliContext([]string{
		"-project-dir=somedir4",
		"-project-pkg=somepkg",
		"-new",
	}))
	require.NoError(t, err)

	require.Equal(t, `Mkdir .typical-tmp
Mkdir cmd/somepkg
Mkdir internal/app
Mkdir app
Create wrapper 'somedir4/typicalw'
`, output.String())
}
