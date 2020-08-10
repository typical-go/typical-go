package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/internal/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func TestApp(t *testing.T) {
	typapp.Name = "some-name"
	typapp.Version = "some-version"
	defer func() {
		typapp.Name = ""
		typapp.Version = ""
	}()
	app := app.App()
	require.Equal(t, "some-name", app.Name)
	require.Equal(t, "some-version", app.Version)
	require.Equal(t, "run", app.Commands[0].Name)
	require.Equal(t, "setup", app.Commands[1].Name)
}
