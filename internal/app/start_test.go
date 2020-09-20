package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/internal/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestApp(t *testing.T) {
	typgo.AppName = "some-name"
	typgo.AppVersion = "some-version"
	defer func() {
		typgo.AppName = ""
		typgo.AppVersion = ""
	}()
	app := app.App()
	require.Equal(t, "some-name", app.Name)
	require.Equal(t, "some-version", app.Version)
	require.Equal(t, "run", app.Commands[0].Name)
	require.Equal(t, "setup", app.Commands[1].Name)
}
