package tmpl

// AppTest template
const AppTest = `package app_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typcore"

	"{{.Pkg}}/app"
)

func TestModule(t *testing.T) {
	a := &app.Module{}
	require.True(t, typcore.IsActionable(a))
	require.True(t, typcore.IsConfigurer(a))
}
`
